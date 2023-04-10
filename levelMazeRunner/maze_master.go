package main

import (
	"maze/common"
)

type MazeMaster struct {
	maze [][]int
	wall int
}

func (mm *MazeMaster) Initialize() {
	// 迷路のルールは・・・
	// ・各数値をステップと呼ぶ
	// ・壁以外のステップは重複不可
	// ・壁以外のステップは連続である必要はない
	// ・最小値＝スタートステップから最大値＝ゴールステップまでを辿る
	// この迷路の場合、1〜9までのステップを辿るルートは4つある
	// [1][2][4][5][7][8][9]
	// [1][2][4][5][6][8][9]
	// [1][2][3][5][7][8][9]
	// [1][2][3][5][6][8][9]
	mm.maze = append(mm.maze, []int{0, 1, 0, 0, 0})
	mm.maze = append(mm.maze, []int{0, 2, 3, 0, 0})
	mm.maze = append(mm.maze, []int{0, 4, 5, 6, 0})
	mm.maze = append(mm.maze, []int{0, 0, 7, 8, 0})
	mm.maze = append(mm.maze, []int{0, 0, 0, 9, 0})
	// 迷路の壁となるステップの値は0
	mm.wall = 0
}

// 与えられた迷路の壁以外の最小値＝スタートステップを検出
func (mm *MazeMaster) GetStart() int {
	filteredArray := common.Filter(
		common.Flat(mm.maze),
		func(item int) bool { return item != mm.wall },
	)
	start := common.Min(filteredArray)
	return start
}

// 与えられた迷路の最大値＝ゴールステップを検出
func (mm *MazeMaster) GetGoal() int {
	filteredArray := common.Filter(
		common.Flat(mm.maze),
		func(item int) bool { return item != 0 },
	)
	end := common.Max(filteredArray)
	return end
}

func main() {
	mm := MazeMaster{}
	mm.Initialize()
	mr := MazeRunner{
		wall:  mm.wall,
		start: mm.GetStart(),
		goal:  mm.GetGoal(),
	}
	mr.Run(mm.maze)
}
