package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println("これはランダムに設定された1～１００までの数字を当てるゲームです")
	rand_num := rand.Intn(99) + 1

	cnt := 0
	for {
		cnt++
		ans := number_guessing_game(rand_num)
		if ans {
			fmt.Printf("試行回数は%d回でした\n", cnt)
			break
		}
	}
}

func number_guessing_game(rand_num int) bool {
	fmt.Printf("適当な数値を入力してください：")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	var str_num string = ""
	for _, c := range input {
		if unicode.IsDigit(c) {
			str_num = str_num + string(c)
			continue
		} else {
			fmt.Println("１～１００の数字を入力してください")
			return false
		}
	}

	usr_ans, _ := strconv.Atoi(str_num)
	if usr_ans < 1 || 100 < usr_ans {
		fmt.Println("１～１００の数字を入力してください")
		return false
	}

	switch {
	case usr_ans < rand_num:
		fmt.Println("もっと大きい数字です")
		return false
	case usr_ans > rand_num:
		fmt.Println("もっと小さい数字です")
		return false
	default:
		fmt.Println("正解！天才！")
		return true
	}

}
