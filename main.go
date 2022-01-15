package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jana-veverkova/math-for-kids-be/games"
)

func main() {
	router := gin.Default()
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.StaticFS("/static", http.Dir("C:\\Users\\Jana\\Documents\\go\\mathForKids"))
	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, games.SimpleEquation())
	})
	router.Run(":8080")
}
