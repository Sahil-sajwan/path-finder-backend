package grid

type Point struct {
	Row int `json:"row"`
	Col int `json:"col"`
}

var Grid [20][20]Point

func check(row int, col int) bool {
	if row >= 0 && row < 20 && col >= 0 && col < 20 {
		return true
	}
	return false
}

func Dfs(p1 Point, p2 Point, visited map[Point]bool, path []Point) []Point {

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	path = append(path, p1)
	visited[p1] = true
	for i := 0; i < 4; i++ {
		row := p1.Row + dx[i]
		col := p1.Col + dy[i]

		if check(row, col) && !visited[Point{row, col}] {
			nxt := Point{
				row,
				col,
			}

			if row == p2.Row && col == p2.Col {
				path = append(path, nxt)
				return path
			}
			res := Dfs(nxt, p2, visited, path)
			if len(res) > 0 {
				return res
			}

		}

	}

	return []Point{}
}
