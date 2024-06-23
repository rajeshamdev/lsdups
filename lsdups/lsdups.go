package lsdups

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type Task struct {
	FilePath string
}

const (
	cksumWorkersCount = 10
)

// map to hold dups. proect by lock for concurrent write.
var dupsMap map[string][]string
var dupsMapMutex sync.Mutex

func Lsdups(dirPath string) {

	if dirPath == "" {
		fmt.Printf("Please select a directory and rerun \n")
		return
	}

	dupsMap = make(map[string][]string)

	tasks := make(chan Task, cksumWorkersCount)
	var cksumWorkerGroup sync.WaitGroup

	// create workers
	for i := 1; i <= cksumWorkersCount; i++ {
		cksumWorkerGroup.Add(1)
		go cksumWorker(i, tasks, &cksumWorkerGroup)
	}

	// create a thread that iterates through dir recursively and sends file names
	// to workers. Workers compute checksum. Finally this thread closes the channel
	// when dir iterate is complete.
	go dirWalker(dirPath, tasks)

	cksumWorkerGroup.Wait()

	fmt.Printf("CPUs : %v\n", runtime.NumCPU())
	fmt.Printf("Goroutines : %v\n", runtime.NumGoroutine())
	fmt.Printf("all tasks processed\n")

	for k, v := range dupsMap {
		// print if there are any dups
		if len(v) > 1 {
			fmt.Printf("%v: %v\n", k, v)
		}
	}
}

func dirWalker(dir string, tasks chan<- Task) {

	walkerCallback := func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return nil
		}

		if !info.IsDir() {
			// send file to worker threads
			tasks <- Task{FilePath: path}
		}
		return nil
	}

	// walks through dir recursively
	filepath.Walk(dir, walkerCallback)

	// close the channel so as workers exit gracefully
	close(tasks)
}

func cksumWorker(id int, tasks <-chan Task, wg *sync.WaitGroup) {

	for task := range tasks {
		cksum, _ := computeCksum(task.FilePath)
		dupsMapMutex.Lock()
		dupsMap[cksum] = append(dupsMap[cksum], task.FilePath)
		dupsMapMutex.Unlock()
		//fmt.Printf("%s, worker: %d, checksum: %s\n", task.FilePath, id, cksum)
	}

	wg.Done()
}

func computeCksum(fname string) (string, error) {

	file, err := os.Open(fname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileHash := sha256.New()

	if _, err := io.Copy(fileHash, file); err != nil {
		return "", err
	}

	hashSum := fileHash.Sum(nil)
	hashString := hex.EncodeToString(hashSum)
	return hashString, nil
}
