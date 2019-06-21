// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
    "time"
    "strings"
)

func main() {
    s, sep := "", ""
    start := time.Now();
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
    secs := time.Since(start).Seconds()
    fmt.Println(secs);

    start = time.Now();
    fmt.Println(strings.Join(os.Args[1:], " "))
    secs = time.Since(start).Seconds()
    fmt.Println(secs);
}
