package cache

import (
	"ru/zakat/L0/internal/models"
	"testing"
)

type repoOrdersStub struct{}

func (r *repoOrdersStub) CreateOrder(order *models.Order) error {
	return nil
}

func (r *repoOrdersStub) FindAll() ([]*models.Order, error) {
	order, err := models.ReadModel("../../model.json")
	if err != nil {
		panic(err)
	}
	orders := []*models.Order{order}
	return orders, nil
}

func (r *repoOrdersStub) ReadModel(filename string) (models.Order, error) {
	return models.Order{}, nil
}

func newCache() *Cache {
	return NewCache(&repoOrdersStub{})
}

func TestFindOrder(t *testing.T) {
	cache := newCache()
	order, ok := cache.FindOrder("b563feb7b2b84b6test")
	if !ok {
		t.Fatalf("wrong ok. want=true, got=false")
	}
	models.TestOrder(t, order)
}

func TestFindNotExistingOrder(t *testing.T) {
	cache := newCache()
	_, ok := cache.FindOrder("123")
	if ok {
		t.Fatalf("wrong ok. want=false, got=true")
	}
}

func TestAddOrder(t *testing.T) {
	cache := newCache()
	cache.AddOrder(&models.Order{
		UID:         "123",
		TrackNumber: "BWIJL",
	})

	order, ok := cache.FindOrder("123")
	if !ok {
		t.Fatalf("wrong ok. want=true, got=false")
	}

	models.TestField(t, "UID", order.UID, "123")
	models.TestField(t, "UID", order.TrackNumber, "BWIJL")
}

func TestFindAll(t *testing.T) {
	cache := newCache()

	orders := cache.FindAll()

	if len(orders) != 1 {
		t.Fatalf("wrong len(orders). want=1, got=%d", len(orders))
	}

	models.TestOrder(t, orders[0])
}
