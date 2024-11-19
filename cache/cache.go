package cache

import (
	"ru/zakat/L0/db"
	"ru/zakat/L0/logger"
	"ru/zakat/L0/models"

	"go.uber.org/zap"
)

type Cache struct {
	orders map[string]*models.Order
	logger *zap.Logger
	repo   *db.Repository
}

func NewCache(repo *db.Repository) *Cache {
	logger := logger.NewLogger()
	orders := repo.FindAll()
	ordersMap := make(map[string]*models.Order)

	for _, order := range orders {
		ordersMap[order.UID] = order
	}

	logger.Info("Cache initialized")
	return &Cache{ordersMap, logger, repo}
}

func (c *Cache) FindAll() []*models.Order {
	ordersArr := make([]*models.Order, 0, len(c.orders))

	for _, order := range c.orders {
		ordersArr = append(ordersArr, order)
	}

	return ordersArr
}

func (c *Cache) FindOrder(uid string) (*models.Order, bool) {
	order, ok := c.orders[uid]
	return order, ok
}

func (c *Cache) AddOrder(order *models.Order) {
	c.repo.CreateOrder(order)
	c.orders[order.UID] = order
}
