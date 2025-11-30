package main

import (
	"gin-quickstart/db"
	"gin-quickstart/handlers"
	"gin-quickstart/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"flag"
)

const (
	ARTICLE_FILE = "/app/jawiki-latest-pages-articles.xml"
	// ARTICLE_FILE = "jawiki-latest-pages-articles.xml"
	MAX_ARTICLES = 4e6
)

var (
	Graph     map[int][]int
	idToTitle map[int]string
	titleToID map[string]int
)

func main() {

	build := flag.Bool("build", false , "")
	flag.Parse()

	err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/api/articles", handlers.SearchArticles)

	go func () {
		if *build {
			idToTitle, titleToID, err = utils.BuildMap(ARTICLE_FILE, MAX_ARTICLES, db.DB, true)
			if err != nil {
				panic(err)
			}
			err = utils.SaveMap(idToTitle, titleToID)
			if err != nil {
				panic(err)
			}

			Graph, err = utils.BuildGraph(ARTICLE_FILE, MAX_ARTICLES, titleToID, db.DB, false)
			if err != nil {
				panic(err)
			}
			err = utils.SaveGraphStreaming(Graph)
			if err != nil {
				panic(err)
			}
		} else {
			idToTitle, titleToID, err = utils.LoadMap()
			if err != nil {
				panic(err)
			}
			Graph, err = utils.LoadGraph()
			if err != nil {
				panic(err)
			}		
		}
		handlers.Graph = Graph
		handlers.IdToTitle = idToTitle
		handlers.TitleToID = titleToID
	}()

	

	r.GET("/api/path", func(c *gin.Context) {
		if handlers.Graph == nil {
			c.JSON(503, gin.H{"error": "Building Graph"})
			return
		}
		handlers.FindShortestPath(c)
	})

	r.Run(":8080")
}
