// NOT WORKING
package main

import (
    "bufio"
    "fmt"
    "log"
    "math"
    "os"
    "sort"
    "strings"
)


func recognize(segments [7]int) int {
    numbers := [10][7]int{
        {1,1,1,0,1,1,1}, // 0
        {0,0,1,0,0,1,0}, // 1
        {1,0,1,1,1,0,1}, // 2
        {1,0,1,1,0,1,1}, // 3
        {0,1,1,1,0,1,0}, // 4
        {1,1,0,1,0,1,1}, // 5
        {1,1,0,1,1,1,1}, // 6
        {1,0,1,0,0,1,0}, // 7
        {1,1,1,1,1,1,1}, // 8
        {1,1,1,1,0,1,1}, // 9
    }

    fmt.Println(segments)
    result := -1
    for i := range numbers {
        if segments == numbers[i] {
            result = i
            break
        }
    }
    return result
}

func test_recognize() {
    test_numbers := [10][7]int{
        {1,1,1,0,1,1,1}, // 0
        {0,0,1,0,0,1,0}, // 1
        {1,0,1,1,1,0,1}, // 2
        {1,0,1,1,0,1,1}, // 3
        {0,1,1,1,0,1,0}, // 4
        {1,1,0,1,0,1,1}, // 5
        {1,1,0,1,1,1,1}, // 6
        {1,0,1,0,0,1,0}, // 7
        {1,1,1,1,1,1,1}, // 8
        {1,1,1,1,0,1,1}, // 9
    }
    for i, number := range test_numbers {
        r := recognize(number)
        if r != i {
            log.Fatal("Bad recognition")
        }
    }
}

func remove_chars(str string, chars string) string {
    r := ""
    for _, s := range str {
        found := false
        for _, c := range chars {
            if s == c {
                found = true
                break
            }
        }
        if found == false {
            r = r + string(s)
        }
    }
    return r
}

func assign_segment(key_segments []string, index int, char_segments []rune,  assignments[]int, previous_chars string) ([]rune, string) {
    number := key_segments[index]
    chars := ""
    if len(previous_chars) > 0 {
        chars = remove_chars(number, previous_chars)
    } else {
        chars = number
    }
    if len(chars) != len(assignments) {
        log.Fatal("String must have # characters: ", len(assignments))
    }
    for i, a := range assignments {
        char_segments[a] = rune(chars[i])
    }
    previous_chars = previous_chars + chars
    return char_segments, previous_chars
}

func assign_segments(digits string) []rune {
    key_segments := strings.Split(digits, " ")
    if (len(key_segments) != 10) {
        log.Fatal("Incorrect number of strings ", key_segments, len(key_segments))
    }
    sort.Slice(key_segments, func(i, j int) bool {
        return len(key_segments[i]) < len(key_segments[j])
    })

    char_segments := make([]rune, 7)
    previous_chars := ""
    char_segments, previous_chars = assign_segment(key_segments, 0, char_segments, []int{2,5}, previous_chars)
    char_segments, previous_chars = assign_segment(key_segments, 1, char_segments, []int{0}, previous_chars)
    char_segments, previous_chars = assign_segment(key_segments, 2, char_segments, []int{1,3}, previous_chars)
    char_segments, previous_chars = assign_segment(key_segments, 9, char_segments, []int{6,4}, previous_chars)

    return char_segments
}

func print_char_segments(char_segments []rune) {
    for _, char := range char_segments {
        fmt.Print(string(char))
    }
    fmt.Println()
}

func recognize_digit(char_segments []rune, digits string) int {
    var digit_segments [7]int

    for _, s := range digits {
        for i, c := range char_segments {
            if s == c {
                digit_segments[i] = 1
                break
            }
        }
    }
    return recognize(digit_segments)
}

func test_recognize_digits() {
    digits := "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab"
    char_segments := assign_segments(digits)

    // other_digits := strings.Split("cdfeb fcadb cdfeb cdbaf", " ")
    other_digits := strings.Split(digits, " ")
    // other_results := [4]int{5,3,5,3}
    other_results := []int{8,5,2,3,7,9,6,4,0,1}
    result := 0.0
    for i := range other_digits {
        d := recognize_digit(char_segments, other_digits[i])
        if d != other_results[i] {
            log.Fatal("Digit must be 5. Instead is ", d)
        }
        fmt.Println(d)
        result = result + math.Pow(10, float64(len(other_digits) - i - 1)) * float64(d)
    }
    fmt.Printf("Answer test: %0.0f\n", result)

}

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // test_recognize()
    // test_recognize_digits()
    // log.Fatal()

    dat, err := os.Open("inputest")
    check(err)

    scanner := bufio.NewScanner(dat)
    total := 0.0
    n := 0
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(n, line)
        segments := strings.Split(line, " | ")
        if len(segments) != 2 {
            log.Fatal("Cannot split line ", line)
        }

        char_segments := assign_segments(segments[0])
        print_char_segments(char_segments)
        digits := strings.Split(segments[1], " ")
        if len(digits) < 1 {
            log.Fatal("Cannot split line digits ", digits)
        }
        result := 0.0
        for i := range digits {
            d := recognize_digit(char_segments, digits[i])
            if d == -1 {
                log.Fatal("Couldn't recognize segment ", digits[i])
            }
            result += math.Pow(10, float64(len(digits) - i - 1)) * float64(d)
        }
        fmt.Printf("%d %0.0f %s\n", n, result, segments[1])
        total += result
        n++
    }
    fmt.Printf("Answer: %0.0f\n", total)
}
