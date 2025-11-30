package handlers

import (
	"gin-quickstart/db"
	"gin-quickstart/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	Graph     map[int][]int
	IdToTitle map[int]string
	TitleToID map[string]int
)

func convertIdToTitle(path []int) []string {
	pathTitles := make([]string, len(path))
	for i, id := range path {
		title, ok := IdToTitle[id]
		if !ok {
			title = ""
		}
		pathTitles[i] = title
	}
	return pathTitles
}

func FindShortestPath(c *gin.Context) {
	startTitle := c.Query("start")
	endTitle := c.Query("end")

	if startTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "start parameter is required",
		})
		return
	}

	if endTitle == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "end parameter is required",
		})
		return
	}
	

	startID, ok := TitleToID[startTitle]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "start article not found"})
		return
	}

	endID, ok := TitleToID[endTitle]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "end article not found"})
		return
	}

	cache := utils.GetCache(startID, endID, db.DB)
	if cache != nil  {
		pathTitles := convertIdToTitle(cache)
		c.JSON(http.StatusOK, gin.H{
			"path":   pathTitles,
			"length": len(cache),
		})
		return
	}

	path, err := utils.BFS(Graph, startID, func(id int) bool {
		return id == endID
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SaveCache(startID,endID, path, db.DB)

	if len(path) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "no path found"})
		return
	}

	pathTitles := convertIdToTitle(path)
	c.JSON(http.StatusOK, gin.H{
		"path":   pathTitles,
		"length": len(path),
	})
}
