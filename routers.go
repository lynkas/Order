package main

import "github.com/gin-gonic/gin"

func routers() *gin.Engine {
	r:=gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api:=r.Group("/api")
	api.GET("/unit", getUnit)
	api.GET("/item", getItem)
	api.GET("/class", getClass)
	api.POST("/order", postOrder)
	api.POST("/register", register)

	admin:=r.Group("/"+Info.Admin)

	admin.GET("/unprocessed", getUnprocessed)
	admin.GET("/findbyphone", findByPhone)
	admin.POST("/addunit", addUnit)
	admin.POST("/additem", addItem)
	admin.POST("/done", orderDone)
	admin.POST("/removeitem", removeItem)
	admin.POST("/removeunit", removeUnit)
	admin.POST("/addclass", addClass)
	admin.POST("/removeclass", removeClass)

	return r
}
