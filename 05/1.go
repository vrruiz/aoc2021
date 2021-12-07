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

func min_max(first int, second int) (int, int) {
    min := first
    max := second
    if first > second {
        min = second
        max = first
    }
    return min, max
}

func main() {
    moves := make([][]int, 0)
    diagram := make([][]int, 0)
    max_x := 0
    max_y := 0

    dat, err := os.Open("input")
    check(err)
    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        line = strings.Replace(line, " -> ", ",", 1)
        values_string := strings.Split(line, ",")
        values := make([]int, 0)
        for i, v := range values_string {
            n, err := strconv.Atoi(v)
            check(err)
            if i % 2 == 0 && n > max_x {
                max_x = n
            } else if i % 2 == 1 && n > max_y {
                max_y = n
            }
            values = append(values, n)
        }
        moves = append(moves, values)
    }
    for x := 0; x <= max_x; x++ {
        row := make([]int, max_y + 1)
        diagram = append(diagram, row)
    }
    // fmt.Println("Max:", max_x, max_y)
    // fmt.Println(moves)

    for i := range moves {
        if moves[i][0] != moves[i][2] && moves[i][1] != moves[i][3] {
            // fmt.Println("Discard:", i, moves[i])
            continue
        }
        x_start, x_end := min_max(moves[i][0], moves[i][2])
        y_start, y_end := min_max(moves[i][1], moves[i][3])
        // fmt.Println("x start, end:", i, x_start, x_end)
        // fmt.Println("y start, end:", i, y_start, y_end)
        for x := x_start; x <= x_end; x++ {
            for y := y_start; y <= y_end; y++ {
                diagram[x][y] = diagram[x][y] + 1
                // fmt.Println("x, y", x, y, diagram[x][y])
            }
        }
        // fmt.Println(moves[i])
    }

    points := 0
    for x := range diagram {
        for y := range diagram[x] {
            if diagram[x][y] >= 2 {
                points++
            }
        }
    }
    fmt.Println("Answer:", points)
}
