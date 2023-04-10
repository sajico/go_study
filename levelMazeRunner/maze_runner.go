package main

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

// 壁、スタート、ゴールの各ステップを保持
type MazeRunner struct {
	wall  int
	start int
	goal  int
}

// mainから呼ばれるメソッド
// ステップの履歴＝ルートを入れるリストと
// スタートステップの座標を引数にして
// move関数を呼び出す
func (mr *MazeRunner) Run(maze [][]int) {
	history := []int{}
	start, err := mr.findStart(maze)
	if err != nil {
		fmt.Println(err)
	} else {
		mr.move(maze, start[0], start[1], history)
	}
}

// 再帰呼出される処理
// x座標＝mazeの1次元目（配列の縦方向）
// y座標＝mazeの2次元目（配列の横方向）
// 一般的なグラフの座標系と異なっているので注意
func (mr *MazeRunner) move(maze [][]int, x int, y int, history []int) {
	// 今回はsliceの初期要素が0だから結局毎回新アドレスになってる？
	// fmt.Printf("history = %p\n", history)
	// fmt.Println(history)

	// 現在座標のステップを取得
	current := maze[x][y]

	if current == mr.goal {
		// もしゴールステップならば
		// ルート履歴に追加して
		// そのルート履歴をコンソールへ出力
		history = append(history, current)
		mr.showHistory(history)
	} else if current != mr.wall && !slices.Contains(history, current) {
		// もしステップが壁でなく、ルート履歴にも含まれていない（＝戻ってない）ならば
		// ルート履歴に追加して
		// 下方向に１つ移動して再帰呼出
		// 右方向に１つ移動して再帰呼出
		// 左方向に１つ移動して再帰呼出
		// 上方向に１つ移動して再帰呼出
		history = append(history, current)
		if x+1 < len(maze) {
			mr.move(maze, x+1, y, mr.cloneHistory(history))
		}
		if y+1 < len(maze[x]) {
			mr.move(maze, x, y+1, mr.cloneHistory(history))
		}
		if y-1 > -1 {
			mr.move(maze, x, y-1, mr.cloneHistory(history))
		}
		if x-1 > -1 {
			mr.move(maze, x-1, y, mr.cloneHistory(history))
		}
	}
}

// スタートステップの座標を取得
// x座標＝mazeの1次元目（配列の縦方向）
// y座標＝mazeの2次元目（配列の横方向）
// 一般的なグラフの座標系と異なっているので注意
func (mr *MazeRunner) findStart(maze [][]int) ([]int, error) {
	for x, row := range maze {
		for y, step := range row {
			if step == mr.start {
				return []int{x, y}, nil
			}
		}
	}
	return []int{}, errors.New("[Error] findStart failed")
}

// ルート履歴のリストを複製する
// 複製しないとどうなるかは試してみましょう
// 今回はsliceの初期要素が0だから結局毎回新アドレスになってる？
func (mr *MazeRunner) cloneHistory(history []int) []int {
	// return history
	return append([]int{}, history...)
}

// ルート履歴をコンソール出力する
func (mr *MazeRunner) showHistory(history []int) {
	for _, item := range history {
		fmt.Print(mr.decorate(item))
	}
	fmt.Println()
}

func (mr *MazeRunner) decorate(item int) string {
	return "[" + strconv.Itoa(item) + "]"
}
