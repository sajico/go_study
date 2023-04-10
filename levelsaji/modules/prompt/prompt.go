package prompt

import (
	"bufio"
	"fmt"
	"levelsaji/modules/types/appmode"
	"levelsaji/modules/types/item"
	"os"
)

type Prompt struct{}

func (prompt *Prompt) scan() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text(), scanner.Err()
}

func (prompt *Prompt) coloredPrintln(color string, value string) {
	fmt.Printf("\033[38;5;"+color+"m%s\033[m\n", value)
}

func (prompt *Prompt) PrintlnGreen(value string) {
	prompt.coloredPrintln("2", value)
}

func (prompt *Prompt) PrintlnBlue(value string) {
	prompt.coloredPrintln("12", value)
}

func (prompt *Prompt) PrintlnOrange(value string) {
	prompt.coloredPrintln("11", value)
}

func (prompt *Prompt) PrintlnRed(value string) {
	prompt.coloredPrintln("1", value)
}

func (prompt *Prompt) selectionNotValid(selection string) {
	fmt.Println()
	prompt.PrintlnRed("[" + selection + "]は有効な選択ではありません")
	fmt.Println()
}

func (prompt *Prompt) PromptModeSelect() (appmode.AppModeType, error) {
	for {
		prompt.PrintlnBlue("モードを選択してください")
		prompt.PrintlnGreen("購入：Enter / 登録(register)：\"r\" / 商品一覧(list)：\"l\" / 終了(quit)：\"q\"")
		mode, err := prompt.scan()
		if err != nil {
			return appmode.Purchase, err
		} else {
			switch mode {
			case appmode.AppModeKeys.Register:
				return appmode.Register, err
			case appmode.AppModeKeys.List:
				return appmode.List, err
			case appmode.AppModeKeys.Purchase:
				return appmode.Purchase, err
			case appmode.AppModeKeys.Quit:
				return appmode.Quit, err
			default:
				prompt.selectionNotValid(mode)
				continue
			}
		}
	}
}

func (prompt *Prompt) PromptRegister(items *item.Items) error {
	prompt.PrintlnBlue("登録する商品の情報を入力してください")
	prompt.PrintlnGreen("商品名,価格,在庫数 のようにカンマ区切りで入力してください")
	for {
		input, err := prompt.scan()
		if err != nil {
			prompt.PrintlnRed("商品登録処理に失敗しました")
			return err
		} else {
			item := item.Item{}
			if err := item.SetPropertiesFromInput(input); err != nil {
				prompt.PrintlnRed(err.Error())
				continue
			} else {
				if err := items.AddItem(item); err != nil {
					prompt.PrintlnRed(err.Error())
					continue
				} else {
					return nil
				}
			}
		}
	}
}
