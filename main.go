package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spetrunin/codingtest/api/plans"
	"github.com/spetrunin/codingtest/api/subscriptions"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/plans", plans.ListHandler)
	r.POST("/subscribe", subscriptions.SubscribeHandler)

	r.Run() // listen and serve on 0.0.0.0:8080
}
