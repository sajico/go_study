package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func expr(input string, index *int) int {
	ret := mul(input, index)

	if input[*index] == '+' {
		*index++
		ret += mul(input, index)
	}
	if input[*index] == '-' {
		*index++
		ret -= mul(input, index)
	}
	return ret
}

func mul(input string, index *int) int {
	ret := primary(input, index)

	if input[*index] == '*' {
		*index++
		ret *= primary(input, index)
	}
	if input[*index] == '/' {
		*index++
		ret /= primary(input, index)
	}

	return ret
}

func primary(input string, index *int) int {
	if input[*index] == '(' {
		*index++
		num := expr(input, index)
		if input[*index] == ')' {
			if *index+1 < len(input) {
				*index++
			}
			return num
		}
	}

	//この辺にエラー処理いれるといい感じかもしれない
	if input[*index] != '(' && input[*index] != ')' && input[*index] != '+' && input[*index] != '-' && input[*index] != '*' && input[*index] != '/' && !unicode.IsDigit(rune(input[*index])) {
		return -1
	}
	// 数字以外のものを最後に指して抜ける
	var num_str string = ""
	for i := *index; i < len(input); i++ {
		if unicode.IsDigit(rune(input[*index])) {
			num_str = num_str + string(input[*index])
			if *index+1 < len(input) {
				*index++
			}
		} else {
			break
		}
	}

	ret, _ := strconv.Atoi(num_str)
	return ret
}

func main() {
	// 読み込み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	}
	input := strings.ReplaceAll(scanner.Text(), " ", "")

	// input := "(2+3)*4"
	index := 0
	// エラー処理考えずに再帰下降構文を実装する
	fmt.Println(expr(input, &index))

}
