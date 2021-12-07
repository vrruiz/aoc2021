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
    fishes := make([]float64, 9)

    dat, err := os.Open("input")
    check(err)
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        values_string := strings.Split(line, ",")
        for _, v := range values_string {
            fmt.Println(v)
            n, err := strconv.Atoi(v)
            check(err)
            fishes[n]++
        }
    }
    fmt.Println("Init", fishes)

    for day := 1; day <= 80; day++ {
        new_fishes := make([]float64, 9)
        for i := 8; i >= 0; i-- {
            if i == 0 {
                new_fishes[8] = fishes[i]
                new_fishes[6] += fishes[i]
            } else {
                new_fishes[i-1] = fishes[i]
            }
        }
        fishes = new_fishes
        var count_fishes float64
        for i := range fishes {
            count_fishes = count_fishes + fishes[i]
        }
        fmt.Printf("%d,%0.0f\n", day, count_fishes)
    }
}
