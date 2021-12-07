package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
    "strconv"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    dat, err := os.Open("input")
    check(err)

    horiz_pos := 0
    depth_pos := 0
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        command := strings.Split(line, " ")
        direction := command[0]
        value, err := strconv.Atoi(command[1])
        check(err)
        switch direction {
            case "up": depth_pos -= value
            case "down": depth_pos += value
            case "forward": horiz_pos += value
            default: log.Fatal("Error")
        }
        fmt.Println(direction, value, "->", horiz_pos, depth_pos)
    }
    fmt.Printf("Answer: %d x %d = %d", horiz_pos, depth_pos, horiz_pos * depth_pos)
}
