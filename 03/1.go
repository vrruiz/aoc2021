package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    dat, err := os.Open("input")
    check(err)

    count_ones := make([]int, 12)
    total := 0
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        for pos, char := range line {
            if char == '1' {
                count_ones[pos]++
            }
        }
        total++
    }

    var gamma_rate int = 0
    var epsilon_rate int = 0
    for i, value := range count_ones {
        shift := 0b1 << (12 - i - 1)
        if value > total / 2 {
            gamma_rate += shift
        } else {
            epsilon_rate += shift
        }
    }
    fmt.Println("Answer: ", gamma_rate * epsilon_rate)
}
