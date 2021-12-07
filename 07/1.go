package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func read_positions(file_name string) []int {
    positions := make([]int, 0)
 
    dat, err := os.Open(file_name)
    check(err)

    scanner := bufio.NewScanner(dat)
    for scanner.Scan() {
        line := scanner.Text()
        values_string := strings.Split(line, ",")
        for _, s := range values_string {
            p, err := strconv.Atoi(s)
            check(err)
            positions = append(positions, p)
        }
    }

    return positions
}

func round(val float64) int {
    if val < 0 {
        return int(val-0.5)
    }
    return int(val+0.5)
}

func max_min_avg(positions []int) (int, int, int, int) {
    max, min, avg, avg_w := 0, 0, 0, 0
    average := 0
    average_w := make(map[int]int)

    for _, pos := range positions {
        if pos > max {
            max = pos
        }
        if pos < min {
            min = pos
        }
        average += pos
        _, ok := average_w[pos]
        if ok == true {
            average_w[pos]++
        } else {
            average_w[pos] = 1
        }
    }
    avg = average / len(positions)
    var avg_f float64 = 0
    for key, value := range average_w {
        avg_f = avg_f + float64(key) * float64(value)
    }
    if len(positions) > 0 {
        avg_w = round(avg_f / float64(len(positions)))
    }
    return max, min, avg, avg_w
}

func abs(v int) int {
    if v > 0 {
        return v
    } else {
        return -v
    }
}

func calculate_cost(positions []int, move_to int) int {
    points := 0
    for i := range positions {
        points += abs(positions[i] - move_to)
    }
    return points
}

func calculate_min_cost(positions []int, from_pos int, to_pos int) (int, int) {
    min_cost := -1
    min_pos := -1
    for pos := from_pos; pos <= to_pos; pos++ {
        cost := calculate_cost(positions, pos)
        if min_cost == -1 || cost < min_cost {
            min_cost = cost
            min_pos = pos
        }
    }
    return min_cost, min_pos
}

func calculate_answer(file_name string) int {
    positions := read_positions(file_name)
    if (len(positions) < 0) {
        log.Fatal("Empty values")
    }

    max_pos, min_pos, avg_pos, avgw_pos := max_min_avg(positions)
    fmt.Printf("min:%d max:%d avg:%d avg_w:%d\n", min_pos, max_pos, avg_pos, avgw_pos)

    cost_min, pos_min := calculate_min_cost(positions, min_pos, max_pos)
    fmt.Printf("cost:%d pos:%d\n", cost_min, pos_min)

    return cost_min
}

func main() {
    answer_test := calculate_answer("inputest")
    fmt.Println("Answer test: ", answer_test)
    if (answer_test != 37) {
        log.Fatal("Answer is not 37")
    }

    answer := calculate_answer("input")
    fmt.Println("Answer: ", answer)
}
