package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rajeshamdev/lsdups/api"
)

var (
	lsdupHTTPServer    *http.Server
	lsdupGoRoutinesCnt int
	lsdupGoRoutinesWG  sync.WaitGroup
)

func lsdupInit() {

	// block all async signals to this server
	signal.Ignore()

	// create buffered signal channel and register below signals:
	//   - SIGINT (Ctrl+C)
	//   - SIGHUP (reload config)
	//   - SIGTERM (graceful shutdown)
	//   - SIGCHLD (handle child processes sending signal to parent) - ignore for now
	api.SigChan = make(chan os.Signal, 1)
	signal.Notify(api.SigChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGCHLD)

	lsdupRouter := gin.Default()
	lsdupRouter.Use(gin.Logger())
	lsdupRouter.GET("/v1/api/dups/list", api.LsdupGet)
	lsdupRouter.GET("/v1/api/dups/health", api.LsdupHealth)
	// lsdupRouter.DELETE("/v1/api/lsdup/delete", api.lsdupDelete)
	lsdupRouter.POST("/v1/api/dups/shutdown", api.LsdupShutdown)

	lsdupHTTPServer = &http.Server{
		Addr:    ":8080",
		Handler: lsdupRouter,
	}
}

func lsdupServerStart() {

	fmt.Printf("lsdupServerStart starting\n")
	err := lsdupHTTPServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Printf("list: %v\n", err)
	}

	lsdupGoRoutinesWG.Done()
}

func signalHandler() {

	for {
		select {

		case sig := <-api.SigChan:

			if sig == syscall.SIGCHLD {
				continue
			} else if sig == syscall.SIGINT || sig == syscall.SIGTERM {
				fmt.Printf("signal received: %v\n", sig)

				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				err := lsdupHTTPServer.Shutdown(ctx)
				if err != nil {
					fmt.Printf("forced shutdown: %v", err)
				} else {
					fmt.Printf("graceful shutdown")
				}
			}
		}
	}

}

func main() {

	lsdupInit()

	go signalHandler()

	lsdupGoRoutinesCnt++
	go lsdupServerStart()
	lsdupGoRoutinesWG.Add(lsdupGoRoutinesCnt)

	lsdupGoRoutinesWG.Wait()

}
