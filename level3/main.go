package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	// inputs := "20"

	var score_str string = ""
	for i, c := range inputs {
		if "0" <= c && c <= "9" {
			score_str = score_str + string(inputs[i])
		} else {
			fmt.Println("数字以外の入力はしないでください")
			return
		}
	}

	score, _ := strconv.Atoi(score_str)

	if score < 0 || 100 < score {
		fmt.Println("0 ～ 100 の間で入力してください")
		return
	}

	switch {
	case 0 <= score && score < 25:
		fmt.Println("不可")
	case 25 <= score && score < 50:
		fmt.Println("可")
	case 50 <= score && score < 75:
		fmt.Println("良")
	case 75 <= score && score < 100:
		fmt.Println("優")
	}
}
