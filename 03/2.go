package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func rating(measures []string, selection []int, column int, most_common bool) []int {
    fmt.Println("=", column)
    if column == len(measures[0]) {
        return selection
    }

    ones := 0
    zeros := 0
    for _, line := range selection {
        if measures[line][column] == '1' {
            ones++
        } else {
            zeros++
        }
    }
    var common byte = '0'
    if most_common == true {
        if ones > zeros {
            common = '1'
        } else if ones == zeros {
            common = '1'
        }
    } else {
        if ones < zeros {
            common = '1'
        }
    }
    selection_new := make([]int, 0)
    for _, line := range selection {
        if measures[line][column] == common {
            selection_new = append(selection_new, line)
        }
    }

    if len(selection_new) == 1 {
        return selection_new
    } else {
        return rating(measures, selection_new, column + 1, most_common)
    }
}

func main() {
    dat, err := os.Open("input")
    check(err)

    var measures = make([]string, 0)
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        measures = append(measures, line)
    }

    selection := make([]int, len(measures))
    for i, _ := range selection {
        selection[i] = i
    }
    oxygen := rating(measures, selection, 0, true)
    co_two := rating(measures, selection, 0, false)

    if len(oxygen) != 1 || len(co_two) != 1 {
        log.Fatal("More than one result")
    }

    oxygen_rating, err := strconv.ParseInt(measures[oxygen[0]], 2, 64)
    check(err)
    co_two_rating, err := strconv.ParseInt(measures[co_two[0]], 2, 64)
    check(err)

    fmt.Println("Answer: ", oxygen_rating * co_two_rating)
}
