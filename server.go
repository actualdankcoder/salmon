package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/", func(c *gin.Context) {
		dat := LoadJSON()
		c.IndentedJSON(200, dat)
	})
	r.POST("/", func(c *gin.Context) {

		var resp SalDat
		c.BindJSON(&resp)
		data := WriteJSON(resp)
		c.IndentedJSON(200, data)
	})
	r.Run("localhost:8406") // listen and serve on localhost:8406
}
