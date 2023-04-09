package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	Name      string
	Price     int
	Inventory int
}

func To_int(str string) (int, bool) {
	var str_num string = ""
	if str[0] == '-' {
		return -1, false
	}
	for _, c := range str {
		if '0' <= c && c <= '9' {
			str_num = str_num + string(str_num)
		} else {
			return -2, false
		}
	}

	ret, _ := strconv.Atoi(str_num)

	return ret, true
}

func is_continue() bool {
	fmt.Println("続けて入力する時は1を終了する場合は0を入力してください")
	for {
		fmt.Printf("1 or 0：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		is_next := scanner.Text()
		if is_next == "1" {
			return true
		} else if is_next == "0" {
			return false
		} else {
			fmt.Println("1か0を入力してください")
			continue
		}
	}
}

func register_mode() map[string]*Item {
	item := make(map[string]*Item)
	fmt.Println("登録モードです")

	var input_name string = ""
	var input_price int = 0
	var input_inventory int = 0
	var flag bool = false

	for {
		fmt.Println("商品名を入力してください")
		fmt.Printf("商品名：")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_name = scanner.Text()

		fmt.Println("価格を入力してください")
		fmt.Printf("価格：")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_price, flag = To_int(scanner.Text())
		if !flag {
			fmt.Println("正の整数を入力してください")
			continue
		}

		fmt.Println("在庫数を入力してください")
		fmt.Printf("在庫数：")
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input_inventory, flag = To_int(scanner.Text())
		if !flag {
			fmt.Println("正の整数を入力してください")
			continue
		}

		item[input_name] = &Item{Name: input_name, Price: input_price, Inventory: input_inventory}

		if is_continue() {
			fmt.Println("次の商品を入力してください")
			continue
		} else {
			break
		}
	}

	return item
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
			item := register_mode()
			fmt.Println(item)
			break
		} else if input[0] == "購入" {
			// customer_mode(item)
			break
		} else {
			fmt.Println("登録か購入を記入してください")
			continue
		}
	}

	//TODO:カートに追加したものの合計金額を出力する
	// 商品名 * 購入数 = xxx￥
}
