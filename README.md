# List Duplicate Files or Objects

**Problem statement**: The idea of building this tool is conceived out of personal necessity. I wanted a tool that provides
insights into files of a directory. I have machines where we duplicated copies from our digital SLR camera. But, it was
very messy and unambiguous to figure out which pics were copied multiple times. We want to be very careful in sorting and then
deleting the pics as they were precious memories of our family. So, I wanted an algorithmically proven way to list duplicates.

This tool iterates through a given directory, finds the SHA-256 hash on each file, and lists the files that have the same
SHA-256 hash. Since SHA-256 hash is unique for file content, this helps to decide the duplicates. The tool is implemented in Golang.
One goroutine iterates through the directory sending the files to 10 worker goroutines. The worker goroutines compute the SHA-256 of
a file and updates the map (which maintains duplicates).

**Compile and usage:**

```
mars ~/ $ git clone https://github.com/rajeshamdev/lsdups.git
mars ~/ $ cd lsdups
mars ~/ $ make build
go build -mod vendor -o dups main.go
mars ~/ $
mars ~/ $ ./dups ls -d .
CPUs : 8
Goroutines : 1
all tasks processed
b1119f3b473e9c3fab3fc10c957603778fdc796b869ce34f3b887074fd5fd943: [.git/refs/heads/main .git/refs/remotes/origin/main]
28b632f327b4b5ab0d2a673149848e3826bf5d0b6afd5ccdda33a7bda54f0190: [.git/logs/HEAD .git/logs/refs/heads/main]
915080be36d4d457414adec82b375a8b3b4160bfe57de31104686d3dfb6d70ef: [README.md README.md.dup]
mars ~/ $ 
```

**Computing sha-256 of a file:**
```
mars ~/ $ shasum -a 256 README.md README.md.dup
915080be36d4d457414adec82b375a8b3b4160bfe57de31104686d3dfb6d70ef  README.md.dup
915080be36d4d457414adec82b375a8b3b4160bfe57de31104686d3dfb6d70ef  README.md
mars ~/ $
```

There is also REST API server written in golang:
```
mars ~/ $ make server
go build -mod vendor -o server server.go
mars ~/ $
mars ~/ $ ./server

Implemented 3 REST APIs:
    1) GET /v1/api/dups/list: list duplicate files (takes dir query param)
    2) GET /v1/api/dups/health: checks the health of the REST API server
    3) POST /v1/api/dups/shutdown: shutdown the REST API server by sending SIGINT signal
    
Testing:
      curl -X GET localhost:8080/v1/api/dups/list?dir="." | python3 -m json.tool
      curl -X GET localhost:8080/v1/api/dups/health | python3 -m json.tool
      curl -X POST localhost:8080/v1/api/dups/shutdown | python3 -m json.tool

```
