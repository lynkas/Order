package main

import "github.com/gin-gonic/gin"

func routers() *gin.Engine {
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/unit", getUnit)
	r.GET("/item", getItem)
	r.POST("/order", postOrder)
	r.POST("/register", register)

	r.GET("/unprocessed", getUnprocessed)
	r.POST("/addunit", addUnit)
	r.POST("/additem", addItem)
	r.POST("/done", orderDone)




	return r
}
