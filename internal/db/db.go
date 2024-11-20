package db

import (
	"encoding/json"
	"fmt"
	"io"
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

type Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewRepository() *Repository {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger()

	return &Repository{db, logger}
}

func (r *Repository) FindAll() []*models.Order {
	var orders []*models.Order

	result := r.db.Preload("Items").Preload("Delivery").Preload("Payment").Find(&orders)

	r.logger.Info(fmt.Sprintf("Find all. Rows affected: %d", result.RowsAffected))

	return orders
}

func (r *Repository) ReadModel(filename string) models.Order {
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
	return order
}

func (r *Repository) CreateOrder(order *models.Order) {
	r.db.Create(order)
}
