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
    for y := 0; y <= max_y; y++ {
        row := make([]int, max_x + 1)
        diagram = append(diagram, row)
    }
    fmt.Println("Max:", max_x, max_y)
    fmt.Println("Dim:", len(diagram[0]), len(diagram))
    // fmt.Println(moves)

    for i := range moves {
        if moves[i][0] != moves[i][2] && moves[i][1] != moves[i][3] {
            // fmt.Println("Discard:", i, moves[i])
            x := moves[i][0]
            x_end := moves[i][2]
            y := moves[i][1]
            y_end := moves[i][3]
            // fmt.Println(x, x_end, y, y_end)
            for x != x_end && y != y_end {
                // fmt.Printf("x:%d y:%d d:%d", x, y, diagram[y][x])
                diagram[y][x] = diagram[y][x] + 1
                if x < x_end {
                    x++
                } else if x > x_end {
                    x--
                }
                if y < y_end {
                    y++
                } else if y > y_end {
                    y--
                }
                // fmt.Printf(" -> x:%d y:%d\n", x, y)
            }
            diagram[y][x] = diagram[y][x] + 1
        } else {
            x_start, x_end := min_max(moves[i][0], moves[i][2])
            y_start, y_end := min_max(moves[i][1], moves[i][3])
            // fmt.Printf("%d x start:%d end:%d\n", i, x_start, x_end)
            // fmt.Printf("%d y start:%d end:%d\n", i, y_start, y_end)
            for y := y_start; y <= y_end; y++ {
                for x := x_start; x <= x_end; x++ {
                        // fmt.Printf("x:%d y:%d (%d x %d)\n", x, y, len(diagram), len(diagram[y]))
                        diagram[y][x] = diagram[y][x] + 1
                }
            }
        }
        fmt.Println(moves[i])
    }

    points := 0
    for y := range diagram {
        for x := range diagram[y] {
            if diagram[y][x] >= 2 {
                points++
            }
            // fmt.Printf("%d", diagram[y][x])
        }
        // fmt.Println()
    }
    fmt.Println("Answer:", points)
}
