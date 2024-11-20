package db

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"ru/zakat/L0/internal/config"
	"ru/zakat/L0/internal/logger"
	"ru/zakat/L0/internal/models"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DbHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort)
)

type RepoOrders interface {
	FindAll() ([]*models.Order, error)
	ReadModel(filename string) (models.Order, error)
	CreateOrder(order *models.Order) error
}

type repoOrders struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewRepository() RepoOrders {
	logger := logger.NewLogger()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to open connection", zap.Error(err))
	}

	return &repoOrders{
		db:  db,
		log: logger,
	}
}

func (r *repoOrders) FindAll() ([]*models.Order, error) {
	var orders []*models.Order

	result := r.db.Preload("Items").Preload("Delivery").Preload("Payment").Find(&orders)

	r.log.Info(fmt.Sprintf("Find all. Rows affected: %d", result.RowsAffected))

	return orders, nil
}

func (r *repoOrders) ReadModel(filename string) (models.Order, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var order models.Order

	json.Unmarshal(byteValue, &order)

	r.db.Create(&order)
	return order, nil
}

func (r *repoOrders) CreateOrder(order *models.Order) error {
	r.db.Create(order)
	return nil
}
