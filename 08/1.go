package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func unique_indexes(segments [10]int) []int {
    var counter [8]int
    for _, v := range segments {
        counter[v] = counter[v] + 1
    }
    unique := make([]int, 0)
    for i, v := range counter {
        if v == 1 {
            unique = append(unique, i)
        }
    }
    unique_i := make([]int, 0)
    for i, v := range segments {
        for _, u := range unique {
            if v == u {
                unique_i = append(unique_i, i)
            }
        }
    }
    return unique_i
}

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func process_input(file_name string) [][][]string {
    dat, err := os.Open(file_name)
    check(err)
    scanner := bufio.NewScanner(dat)
    input_digits := make([][][]string, 0)
    for scanner.Scan() {
        line := scanner.Text()
        s := strings.Split(line, " | ")
        words := make([][]string, 0)
        for _, ss := range s {
            w := strings.Split(ss, " ")
            words = append(words, w)
        }
        input_digits = append(input_digits, words)
    }
    return input_digits
}

func get_digits(digits [][][]string, index int) [][]string {
    g_digits := make([][]string, 0)
    for i := range digits {
        g_digits = append(g_digits, digits[i][index])
    }
    return g_digits
}

func count_unique_digits(digits [][]string, segments [10]int, unique []int) int {
    count := 0
    for i := range digits {
        for n := range digits[i] {
            s := digits[i][n]
            for _, u := range unique {
                if len(s) == segments[u] {
                    fmt.Print(s, " ", len(s), " ")
                    count++
                    break
                }
            }
        }
        fmt.Println()
    }
    return count
}

func count_digits(file_name string, segments [10]int, unique []int) int {
    digits := process_input(file_name)
    // fmt.Println(digits)    
    output_digits := get_digits(digits, 1)
    // fmt.Println(output_digits)
    count := count_unique_digits(output_digits, segments, unique)
    return count
}

func main() {
    segments := [10]int{6,2,5,5,4,5,6,3,7,6}
    unique := unique_indexes(segments)
    fmt.Println(unique)

    count := count_digits("inputest", segments, unique)
    fmt.Println("Answer test:", count)
    if count != 26 {
        log.Fatal("Answer not 26")
    }

    count = count_digits("input", segments, unique)
    fmt.Println("Answer:", count)
}