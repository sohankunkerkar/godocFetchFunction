package main

import (
        "flag"
        "log"
        "net/http"
        "os"
        "github.com/davecgh/go-spew/spew"
        "github.com/golang/gddo/doc"
        "fmt"
        "bufio"
        "os/exec"
        "github.com/golang/gddo/gosrc"
)

var (
        etag  = flag.String("etag", "", "Etag")
        local = flag.Bool("local", false, "Get package from local directory.")
)

func main() {
        flag.Parse()
        if len(flag.Args()) != 1 {
                log.Fatal("Usage: go run print.go importPath")
        }
        path := flag.Args()[0]

        var (
                pdoc *doc.Package
                err  error
        )
        if *local {
                gosrc.SetLocalDevMode(os.Getenv("GOPATH"))
        }
        pdoc, err = doc.Get(http.DefaultClient, path, *etag)
        if err != nil {
                log.Fatal(err)
        }
       str := spew.Sdump(pdoc)
       fileHandle, _ := os.Create("output.txt")
        writer := bufio.NewWriter(fileHandle)
        defer fileHandle.Close()
        fmt.Fprintln(writer,str)
        writer.Flush()
       ls := exec.Command("/bin/bash","golang.sh")
       op, err := ls.Output()
        if err != nil {
                fmt.Println(err)
                return
        }
        fmt.Printf("Command ran: %s\n", string(op))
    }