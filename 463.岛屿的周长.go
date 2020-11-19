/*
 * @lc app=leetcode.cn id=463 lang=golang
 *
 * [463] 岛屿的周长
 */
package main

import "fmt"

func main() {
	fmt.Println(islandPerimeter([][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}))
}

// 周长 = 4 * 土地 - 2 * 接临的边
// 采用深度遍历的方式，如果正下方或右方是土地，则计算有一条接临的边
// @lc code=start
func islandPerimeter(grid [][]int) int {
	land, border := 0, 0
	height, width := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				land++
				if i < height-1 && grid[i+1][j] == 1 {
					border++
				}
				if j < width-1 && grid[i][j+1] == 1 {
					border++
				}
			}

		}
	}
	return land*4 - 2*border
}

// @lc code=end
