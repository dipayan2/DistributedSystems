
package main

import (
    // "fmt"
    "os/exec"
    "sync"
    "bufio"
    "os"
    "io"
    // "sync"
    // "strings"
)

func ExecuteCommand(comd string, outputfile string, wg *sync.WaitGroup) {
    // split_command := strings.Split(comd, " ")



    // cmd := exec.Command("echo", "'WHAT THE HECK IS UP'")
    cmd := exec.Command("sh","-c",comd)
    // open the out file for writing
    outfile, err := os.Create(outputfile)
    if err != nil {
        panic(err)
    }
    defer outfile.Close()

    stdoutPipe, err := cmd.StdoutPipe()
    if err != nil {
        panic(err)
    }

    writer := bufio.NewWriter(outfile)
    defer writer.Flush()

    err = cmd.Start()
    if err != nil {
        panic(err)
    }

    go io.Copy(writer, stdoutPipe)
    cmd.Wait()
    wg.Done()
}


func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)
    go ExecuteCommand("ls","test.txt", wg)
    go ExecuteCommand("ls -ltr", "out.txt",wg)
    wg.Wait()

}