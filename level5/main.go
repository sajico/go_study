package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	Name      string
	Price     int
	Inventory int
}

func register_mode() {
	item := make(map[string]*Item)

}

func main() {
	for {
		fmt.Printf("登録or購入：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")

		if len(input) != 1 {
			fmt.Println("登録か購入を記入してください")
			continue
		}

		if input[0] == "登録" {
			register_mode()
			break
		} else if input[0] == "購入" {
			customer_mode()
			break
		} else {
			fmt.Println("登録か購入を記入してください")
			continue
		}
	}

	//TODO:カートに追加したものの合計金額を出力する
	// 商品名 * 購入数 = xxx￥
}
