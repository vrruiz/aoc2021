package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func adyacent(board [][]int, i int, j int) []int {
	if i >= len(board[0]) || j >= len(board) {
		fmt.Println(i, len(board[0]), j, len(board))
		log.Fatal("adyacent: i, j bigger than board dimensions")
	}
	values := make([]int, 0)
	pos := [][]int{{-1, 0}, {0, 0}, {1, 0}, {0, -1}, {0, 1}}
	for n := range pos {
		x := pos[n][0] + i
		y := pos[n][1] + j
		if x >= 0 && y >= 0 && y < len(board) && x < len(board[0]) {
			v := board[y][x]
			values = append(values, v)
		}
	}
	sort.Slice(values, func(p, q int) bool {
		return values[p] < values[q]
	})
	return values
}

func minimum_board(board [][]int) [][]int {
	min_board := make([][]int, 0)
	for y := range board {
		min_row := make([]int, len(board[y]))
		for x := range board[y] {
			v := board[y][x]
			ady := adyacent(board, x, y)
			if ady[0] == v && ady[1] != v {
				// fmt.Println("Min value", x, y, v)
				min_row[x] = 1
			}
		}
		min_board = append(min_board, min_row)
	}
	return min_board
}

func read_board(file_name string) [][]int {
	dat, err := os.Open(file_name)
	check(err)

	board := make([][]int, 0)
	scanner := bufio.NewScanner(dat)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, len(line))
		for i, r := range line {
			s := string(r)
			row[i], err = strconv.Atoi(s)
			check(err)
		}
		if len(board) > 1 && len(row) != len(board[0]) {
			log.Fatal("read_board: All rows must have the same size")
		}
		board = append(board, row)
	}
	return board
}

func sum_board(board [][]int, min_board [][]int) int {
	if len(board) != len(min_board) && len(board[0]) != len(min_board[0]) {
		log.Fatal("sum_boards: Size of boards don't match")
	}
	total := 0
	for y := range board {
		for x := range board[y] {
			// fmt.Println(x, y)
			if min_board[y][x] == 1 {
				total += board[y][x] + 1
			}
		}
	}
	return total
}

func test() {
	board := read_board("inputest")
	if len(board) != 5 && len(board[0]) != 9 {
		log.Fatal("test: Incorrect size of the board")
	}
	// fmt.Println(board)
	min_board := minimum_board(board)
	// fmt.Println(min_board)
	sum := sum_board(board, min_board)
	if sum != 15 {
		log.Fatal("test: Sum must be 15")
	}
}

func run() {
	board := read_board("input")
	if len(board) != 100 && len(board[0]) != 100 {
		log.Fatal("run: Incorrect size of the board")
	}
	// fmt.Println(board)
	min_board := minimum_board(board)
	// fmt.Println(min_board)
	sum := sum_board(board, min_board)
	fmt.Println("Answer:", sum)
}

func main() {
	// test()
	run()
}
