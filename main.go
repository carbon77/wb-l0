package main

import (
	"fmt"
	"net/http"
	"ru/zakat/L0/cache"
	"ru/zakat/L0/db"
	"ru/zakat/L0/env"
	"ru/zakat/L0/kafka"
	"ru/zakat/L0/logger"
	"ru/zakat/L0/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	log := logger.NewLogger()
	repo := db.NewRepository()
	cache := cache.NewCache(repo)
	orderProducer := kafka.NewProducer()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	go kafka.NewConsumer(cache)

	r.GET("/orders/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		order, ok := cache.FindOrder(uid)
		if !ok {
			c.HTML(http.StatusNotFound, "notfound.html", gin.H{})
			return
		}

		c.HTML(http.StatusOK, "order.html", order)
	})

	api := r.Group("/api")
	{
		api.POST("/orders/send", func(c *gin.Context) {
			var order models.Order

			if err := c.BindJSON(&order); err != nil {
				log.Error("Failed to read request body", zap.Error(err))
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Failed to read request body",
				})
				return
			}

			orderProducer.SendOrder("orders", &order)
			c.JSON(http.StatusOK, gin.H{})
		})

		api.GET("/orders", func(c *gin.Context) {
			orders := cache.FindAll()
			c.JSON(http.StatusOK, orders)
		})

		api.GET("/orders/:uid", func(c *gin.Context) {
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
	}

	port := env.GetEnv("L0_PORT", "8080")
	host := env.GetEnv("L0_HOST", "localhost")
	log.Info("Starting server...", zap.String("port", port))
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
