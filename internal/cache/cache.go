package cache

import (
	"ru/zakat/L0/internal/db"
	"ru/zakat/L0/internal/models"
	"sync"
)

type ICache interface {
	FindAll() []*models.Order
	AddOrder(order *models.Order)
	FindOrder(uid string) (*models.Order, bool)
}

type Cache struct {
	orders map[string]*models.Order
	repo   db.RepoOrders
	mu     *sync.Mutex
}

func NewCache(repo db.RepoOrders) *Cache {
	orders, err := repo.FindAll()
	if err != nil {
		panic(err)
	}
	ordersMap := make(map[string]*models.Order)

	for _, order := range orders {
		ordersMap[order.UID] = order
	}

	return &Cache{ordersMap, repo, &sync.Mutex{}}
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
	c.mu.Lock()
	err := c.repo.CreateOrder(order)
	if err != nil {
		panic(err)
	}
	c.orders[order.UID] = order
	c.mu.Unlock()
}
