package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
    dat, err := os.Open("input")
    check(err)
    defer dat.Close()

    scanner := bufio.NewScanner(dat)
    i := 0
    prev_value := 0
    bigger := 0
    for scanner.Scan() {
        line := scanner.Text()
        value, err := strconv.Atoi(line)
        check(err)
        if i > 0 && value > prev_value {
            bigger++
            fmt.Printf("%d > %d (%d)\n", value, prev_value, bigger)
        }
        prev_value = value
        i++
    }
    fmt.Printf("Answer: %d\n", bigger)

    err = scanner.Err()
    check(err)
}
