# lsdups - List Duplicate Files or Objects

**Problem statement**: The idea of building this tool is conceived out of personal necessity. I wanted a tool that provides
insights into files of a directory. I have machines where we duplicated copies from our digital SLR camera. But, it was
very messy and unambiguous to figure out which pics were copied multiple times. We want to be very careful in sorting and then
deleting the pics as they were precious memories of our family. So, I wanted an algorithmically proven way to list duplicates.

This tool iterates through a given directory, finds the SHA-256 hash on each file, and lists the files that have the same
SHA-256 hash. Since SHA-256 hash is unique for file content, this helps to decide the duplicates. The tool is implemented in Golang.
One goroutine iterates through the directory sending the files to 10 worker goroutines. The worker threads compute the SHA-256 of a file
and updates the map (which maintains duplicates).

```
git clone https://github.com/rajeshamdev/lsdups.git
go build -mod=vendor -o lsdups main.go
lsdups ls (this picks up the files from current working dir)
```


