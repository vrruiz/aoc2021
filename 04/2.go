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

func print_board(board [][]int, score [][]bool) {
    for r := range score {
        for c := range score[r] {
            if score[r][c] == true {
                fmt.Printf("%3d!", board[r][c])
            } else {
                fmt.Printf("%4d", board[r][c])
            }
        }
        fmt.Println()
    }
    fmt.Println()
}

func check_board(board [][]int, score [][]bool, row int, column int) bool {
    win := true
    for c := range score[row] {
        if score[row][c] == false {
            win = false
            break
        }
    }
    if win == true {
        return win
    }
    win = true
    for r := 0; r < 5; r++ {
        if score[r][column] == false {
            win = false
            break
        }
    }
    return win
}

func sum_unmarked_board(board [][]int, score [][]bool) int {
    sum := 0
    for r := range score {
        for c, v := range score[r] {
            if v == false {
                sum += board[r][c]
            }
        }
    }
    return sum
}

func contains(values []int, value int) bool {
    for _, v := range values {
        if v == value {
            return true
        }
    }
    return false
}

func main() {
    numbers := make([]int, 0)
    boards := make([][][]int, 0)

    dat, err := os.Open("input")
    check(err)

    scanner := bufio.NewScanner(dat)
    var current_board [][]int;
    i := 0
    for scanner.Scan() {
        line := scanner.Text()
        if i == 0 {
            number_list := strings.Split(line, ",")
            for _, c := range number_list {
                n, err := strconv.Atoi(c)
                check(err)
                numbers = append(numbers, n)
            }
        } else if line != "" {
            line_list_string := strings.Fields(line)
            if len(line_list_string) != 5 {
                log.Fatal(line)
            }
            line_list := make([]int, 0)
            for _, c := range line_list_string {
                if c != "" {
                    n, err := strconv.Atoi(c)
                    check(err)
                    line_list = append(line_list, n)
                }
            }
            current_board = append(current_board, line_list)
            if len(current_board) == 5 {
                boards = append(boards, current_board)
                current_board = make([][]int, 0)
            }
        }
        i++
    }

    scores := make([][][]bool, len(boards))
    for i := range scores {
        scores[i] = make([][]bool, 5)
        for n := range scores[i] {
            scores[i][n] = make([]bool, 5)
        }
    }

    fmt.Println(numbers)
    win_boards := make([]int, 0)
    value := 0
    loop:
    for i, n := range numbers {
        fmt.Printf("%03d. %d\n", i, n)
        for b := range boards {
            for r := range boards[b] {
                for c := range boards[b][r] {
                    if boards[b][r][c] == n {
                        scores[b][r][c] = true
                        if contains(win_boards, b) == false {
                            win := check_board(boards[b], scores[b], r, c)
                            if win == true {
                                value = n
                                win_boards = append(win_boards, b)
                                if len(win_boards) == len(boards) {
                                    break loop
                                }
                            }
                        }
                    }
                }
            }
        }
    }
    last_board := win_boards[len(win_boards) - 1]
    fmt.Println(last_board, value)
    sum := sum_unmarked_board(boards[last_board], scores[last_board])
    print_board(boards[last_board], scores[last_board])
    fmt.Println("Answer: ", sum * value, sum, value)

}
