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
    // for _, el := range perms_join {
        // fmt.Println(el)
    // }
    // fmt.Printf("get indices of %s\n", substr)
    // fmt.Println(indice_list)
    // fmt.Println(left_win_displacement)
    // get slice of indice_list[i] - (substr_len * word_count) - substr_len
    // fmt.Println(right_win_displacement)
    // don't need the iterator here
    // fmt.Println(s)
    // fmt.Println("-----------loop----------------")
        // for each window, look for each permutation
    // get slice of indice_list[i] - (substr_len * word_count) - substr_len
        // fmt.Println("begin slide")
            // fmt.Printf("iter n: %d\n", i)
                // fmt.Println("skip left")
                // fmt.Println("skip right")
            // fmt.Printf("str in window: %s\n", str)
                    // fmt.Printf("Found substr: %s at %d\n", perm, index)
                    // slide forward

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
    indice_list := make([]int, 0)
    start_index := 0
    substr := words[0]
    substr_len := len(words[0])
    n_of_words := len(words)
    s_len := len(s)
    
    for {
        index := strings.Index(s[start_index:], substr)
        if index == -1 {
            break
        }
        index = index + start_index

        indice_list = append(indice_list, index)

        start_index = index + substr_len
    }
    left_win_displacement := (substr_len * n_of_words) - substr_len
    right_win_displacement := substr_len
    result := make([]int, 0)
    var visited map[int]bool
    visited = make(map[int]bool, s_len)
    for _, index := range indice_list {
        initial_left := index - left_win_displacement
        initial_right := index + right_win_displacement
        for i := 0; i < (n_of_words); i++ {
            step := (substr_len * i)
            n_step_l := initial_left + step
            if visited[n_step_l] == true {
                continue
            }
            n_step_r := initial_right + step
            if n_step_l < 0 {
                visited[n_step_l] = true
                continue
            }
            if n_step_r > s_len {
                continue
            }
            visited[n_step_l] = true
            str := s[initial_left + step:initial_right + step]
            
            for _, perm := range perms_join {
                if str == perm {
                    result = append(result, initial_left + step)
                    break
                }
            }
        }
    }
    return result
}


func main() {
    s1 := "qweqweABCDqweqweCDABqweqweACDBgood"
    words1 := []string{"AB", "CD"}
    // perms := generatePermutations(words1)
    // for _, p := range perms {
    //     fmt.Println(p)
    // }
    fmt.Println(findSubstring(s1, words1))
    s2:= "barfoothefoobarman"
    words2 := []string{"foo","bar"}
    fmt.Println(findSubstring(s2, words2))

    s3 := "wordgoodgoodgoodbestword"
    words3 := []string{"word","good","best","word"}
    fmt.Println(findSubstring(s3, words3))

    s4 := "barfoofoobarthefoobarman"
    words4 := []string{"bar","foo","the"}
    fmt.Println(findSubstring(s4, words4))

    s5 := "wordgoodgoodgoodbestword"
    words5 := []string{"word","good","best","good"}
    fmt.Println(findSubstring(s5, words5))

    s6 := "a"
    words6 := []string{"a"}
    fmt.Println(findSubstring(s6, words6))

    s7 := "aaaaaaaaaaaaaa"
    words7 := []string{"aa","aa"}

    // [0,2,4,6,8,10] (current)
    // [0,1,2,3,4,5,6,7,8,9,10] (expected)
    fmt.Println(findSubstring(s7, words7))
}
