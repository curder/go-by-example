package main

import (
	"fmt"
	"os"
)

type point struct {
	i int
	j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func main() {
	var (
		maze [][]int
		row  []int
		r    int
	)
	maze = readMaze("src/others/cases/maze/maze.in")
	//for _, row = range maze {
	//	for _, r = range row {
	//		fmt.Printf("%d ", r)
	//	}
	//	fmt.Println()
	//}

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row = range steps {
		for _, r = range row {
			fmt.Printf("%3d ", r)
		}
		fmt.Println()
	}
}

func walk(maze [][]int, start, end point) [][]int {
	var (
		steps   [][]int // 行进的路径
		index   int
		q       []point
		current point // 当前节点
		next    point // 下一个节点
		val     int
		ok      bool
	)

	steps = make([][]int, len(maze))
	for index = range steps {
		steps[index] = make([]int, len(maze[index]))
	}

	q = []point{start}
	for len(q) > 0 {
		current = q[0]
		q = q[1:]

		if current == end { // 到达终点
			break
		}

		for _, dir := range dirs {
			next = current.add(dir)

			// 下一个节点需要是0才能探索，并且未到达过
			if val, ok = next.at(maze); !ok || val == 1 {
				continue
			}
			if val, ok = next.at(steps); !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			at, _ := current.at(steps)
			steps[next.i][next.j] = at + 1

			q = append(q, next)
		}
	}
	return steps
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) { // 上下都越界了
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func (p point) add(r point) point {
	return point{
		i: p.i + r.i,
		j: p.j + r.j,
	}
}

//
func readMaze(fileName string) [][]int {
	var (
		maze [][]int
		row  int
		col  int
		file *os.File
		err  error
	)
	if file, err = os.Open(fileName); err != nil {
		panic(err)
	}

	_, err = fmt.Fscanf(file, "%d %d", &row, &col)
	maze = make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			_, err = fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}
