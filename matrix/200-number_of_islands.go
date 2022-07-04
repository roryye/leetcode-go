package matrix

func numIslands(grid [][]byte) int {
	var res int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				res++
				bfs(&grid, i, j)
			}
		}
	}
	return res
}

type Coor struct {
	x int
	y int
}

var dx = [...]int{-1, 0, 1, 0}
var dy = [...]int{0, 1, 0, -1}

func bfs(grid *[][]byte, x, y int) {
	g := *grid
	var queue []Coor
	queue = append(queue, Coor{x, y})
	(*grid)[x][y] = '2'

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {
			newX := temp.x + dx[i]
			newY := temp.y + dy[i]
			if newX >= 0 && newX < len(g) && newY >= 0 && newY < len(g[0]) && g[newX][newY] == '1' {
				queue = append(queue, Coor{newX, newY})
				g[newX][newY] = '2'
			}
		}
	}
}
