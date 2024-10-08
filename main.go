package main

import (
	"pathfinder/grid"
	"pathfinder/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(middleware.OptionsMiddleware())
	r.POST("/find-path", func(c *gin.Context) {

		row1, _ := strconv.Atoi(c.PostForm("rowP1"))
		col1, _ := strconv.Atoi(c.PostForm("colP1"))
		row2, _ := strconv.Atoi(c.PostForm("rowP2"))
		col2, _ := strconv.Atoi(c.PostForm("colP2"))
		p1 := grid.Point{
			Row: row1,
			Col: col1,
		}
		p2 := grid.Point{
			Row: row2,
			Col: col2,
		}
		path := make([]grid.Point, 0)
		visited := make(map[grid.Point]bool)
		for i := 0; i < 20; i++ {
			for j := 0; j < 20; j++ {
				visited[grid.Point{Row: i, Col: j}] = false
			}
		}
		path = grid.Dfs(p1, p2, visited, path)
		c.JSON(200, path)
	})

	r.Run(":8000")

}
