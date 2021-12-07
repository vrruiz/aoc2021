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
    // Variable definition
    queue := make([]int, 0)
    i := 0
    bigger := 0
    window_prev := 0

    // Open file
    dat, err := os.Open("input")
    check(err)
    defer dat.Close()

    // Read values
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        value, err := strconv.Atoi(line)
        check(err)
        if i >= 3 {
            window := window_prev - queue[0] + value
            queue = queue[1:]
            queue = append(queue, value)
            if (window > window_prev) {
                bigger++
            }
            fmt.Printf("%d %d %d\n", window_prev, window, bigger)
            window_prev = window
        } else {
            queue = append(queue, value)
            window_prev += value
            fmt.Printf("%d\n", window_prev)
        }
        i++
    }
    fmt.Printf("Answer: %d\n", bigger)

    err = scanner.Err()
    check(err)
}
