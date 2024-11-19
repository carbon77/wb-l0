package main

import (
	"fmt"
	"net/http"
	"ru/zakat/L0/cache"
	"ru/zakat/L0/db"
	"ru/zakat/L0/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := logger.NewLogger()
	repo := db.NewRepository()
	cache := cache.NewCache(repo)
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world",
		})
	})

	r.GET("/orders", func(c *gin.Context) {
		orders := cache.FindAll()
		c.JSON(http.StatusOK, orders)
	})

	r.GET("/orders/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		order, ok := cache.FindOrder(uid)

		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"message": fmt.Sprintf("Order not found with id=%s", uid),
			})
			return
		}

		c.JSON(http.StatusOK, order)
	})

	logger.Info("Starting server...")
	r.Run()
}
