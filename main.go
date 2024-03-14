package main

import (
	"fmt"
	"strings"
)

func abs(src int) int {
    if src < 0 {
        return -src
    }
    return src
}

func divide(dividend, divisor int) int {
    res_sign := true
    if dividend <= -2147483648 && divisor == -1 {
        return 2147483647 
    }

    if dividend > 0 && divisor > 0 {
        res_sign = true
    } else if dividend < 0 && divisor < 0 {
        divisor = abs(divisor)
        dividend = abs(dividend)
    } else if dividend < 0 {
        res_sign = false
        dividend = abs(dividend)
    } else if divisor < 0 {
        res_sign = false
        divisor = abs(divisor)
    }

    result := 0
    for dividend >= divisor {
        tmp := divisor
        mul := 1
        for tmp << 1 <= dividend {
            tmp = tmp << 1
            mul = mul << 1
        }
        dividend -= tmp
        result += mul
    }

    if res_sign == false {
        return -result
    }
    return result
}

func testdivide() {
    fmt.Printf("10/3: %d\n", divide(10, 3))
    fmt.Printf("6/2: %d\n", divide(6, 2))
    fmt.Printf("6/3: %d\n", divide(6, 3))
    fmt.Printf("14/2: %d\n", divide(14, 2))
    fmt.Printf("7/2: %d\n", divide(7, 2))
    fmt.Printf("7/3: %d\n", divide(7, 3))

    fmt.Printf("10/-3: %d\n", divide(10, -3))
    fmt.Printf("6/-2: %d\n", divide(6, -2))
    fmt.Printf("6/-3: %d\n", divide(6, -3))
    fmt.Printf("14/-2: %d\n", divide(14, -2))
    fmt.Printf("7/-2: %d\n", divide(7, -2))
    fmt.Printf("7/-3: %d\n", divide(7, -3))

    fmt.Printf("-10/3: %d\n", divide(-10, 3))
    fmt.Printf("-6/2: %d\n", divide(-6, 2))
    fmt.Printf("-6/3: %d\n", divide(-6, 3))
    fmt.Printf("-14/2: %d\n", divide(-14, 2))
    fmt.Printf("-7/2: %d\n", divide(-7, 2))
    fmt.Printf("-7/3: %d\n", divide(-7, 3))

    fmt.Printf("-10/-3: %d\n", divide(-10, -3))
    fmt.Printf("-6/-2: %d\n", divide(-6, -2))
    fmt.Printf("-6/-3: %d\n", divide(-6, -3))
    fmt.Printf("-14/-2: %d\n", divide(-14, -2))
    fmt.Printf("-7/-2: %d\n", divide(-7, -2))
    fmt.Printf("-7/-3: %d\n", divide(-7, -3))

    fmt.Printf("int_max/1: %d\n", divide(2147483648, 1))
    fmt.Printf("int_max/1: %d\n", divide(2147483648, -1))
    fmt.Printf("int_max/1: %d\n", divide(-2147483648, 1))
    fmt.Printf("int_max/1: %d\n", divide(-2147483648, -1))
}

func generatePermutations(words []string) [][]string {
    if (len(words) == 1) {
        return [][]string{words}
    }

    var result [][]string
    for i, elem := range words {
        remaining := make([]string, len(words) - 1)
        copy(remaining[:i], words[:i])
        copy(remaining[i:], words[i+1:])

        sub_permutations := generatePermutations(remaining)

		for _, perm := range sub_permutations {
			permWithElem := append([]string{elem}, perm...)
			result = append(result, permWithElem)
		}
    }
    return result
}

func findSubstring(s string, words []string) []int {
    perms_separate := generatePermutations(words)
    
    perms_join := make([]string, len(perms_separate))
    for i, el := range perms_separate {
        perms_join[i] = strings.Join(el, "")
    }
    for _, el := range perms_join {
        fmt.Println(el)
    }
    indice_list := make([]int, 0)
    start_index := 0
    substr := words[1]
    substr_len := len(words[1])
    n_of_words := len(words)

    fmt.Printf("get indices of %s\n", substr)
    
    for {
        index := strings.Index(s[start_index:], substr)
        if index == -1 {
            break
        }
        index = index + start_index

        indice_list = append(indice_list, index)

        start_index = index + substr_len
    }
    fmt.Println(indice_list)

    // get slice of indice_list[i] - (substr_len * word_count) - substr_len
    // get slice of indice_list[i] - (substr_len * word_count) - substr_len
    left_win_displacement := (substr_len * n_of_words) - substr_len
    right_win_displacement := substr_len
    fmt.Println(left_win_displacement)
    fmt.Println(right_win_displacement)
    // don't need the iterator here
    result := make([]int, 0)
    fmt.Println(s)
    for _, index := range indice_list {
        // for each window, look for each permutation
        initial_left := index - left_win_displacement
        initial_right := index + right_win_displacement
        fmt.Println("begin slide")
        for i := 0; i < (n_of_words); i++ {
            step := (substr_len * i)
            str := s[initial_left + step:initial_right + step]
            fmt.Printf("str in window: %s\n", str)
            for _, perm := range perms_join {
                if str == perm {
                    fmt.Printf("Found substr: %s at %d\n", perm, index)
                    result = append(result, initial_left + step)
                    // slide forward
                }
            }
        }
    }
    fmt.Println(result)
    return []int{0, 0}
}


func main() {
    s := "qweqweABCDqweqweCDABqweqweACDBgood"
    words1 := []string{"AB", "CD"}
//    words2 := []string{"qwe", "hghg", "asd"}
//    words3 := []string{"word","good","best","word"}
    perms := generatePermutations(words1)
    for _, p := range perms {
        fmt.Println(p)
    }
//     perms = generatePermutations(words2)
//     for _, p := range perms {
//         fmt.Println(p)
//     }
//     perms = generatePermutations(words3)
//     for _, p := range perms {
//         fmt.Println(p)
//     }
    findSubstring(s, words1)
//    findSubstring(s, words2)
//    findSubstring(s, words3)
}
