package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
)

func main() {
        reader := bufio.NewReader(os.Stdin)
        for {
                fmt.Print("Shell > ")
                text, _ := reader.ReadString('\n')
                text = text[:len(text)-1] // Remove newline

                cmd := exec.Command("bash", "-c", text)
                cmd.Stdout = os.Stdout
                cmd.Stderr = os.Stderr
                err := cmd.Run()
                if err != nil {
                        fmt.Println(err)
                }
        }
}