package models

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"
	"time"
)

type Order struct {
	UID               string    `json:"order_uid" gorm:"primaryKey;column:order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`

	Delivery Delivery `json:"delivery" gorm:"foreignKey:OrderUID;contraint:OnDelete:CASCADE;"`
	Payment  Payment  `json:"payment" gorm:"foreignKey:OrderUID;contraint:OnDelete:CASCADE;"`
	Items    []Item   `json:"items" gorm:"foreignKey:OrderUID;contraint:OnDelete:CASCADE;"`
}

type Delivery struct {
	ID       int64  `gorm:"primaryKey;column:delivery_id"`
	OrderUID string `gorm:"column:order_uid"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Zip      string `json:"zip"`
	City     string `json:"city"`
	Address  string `json:"address"`
	Region   string `json:"region"`
	Email    string `json:"email"`
}

type Payment struct {
	ID           int64   `gorm:"primaryKey;column:payment_id"`
	OrderUID     string  `gorm:"column:order_uid"`
	Transaction  string  `json:"transaction"`
	RequestId    string  `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       float32 `json:"amount"`
	PaymentDt    int64   `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost float64 `json:"delivery_cost"`
	GoodsTotal   float64 `json:"goods_total"`
	CustomFee    float64 `json:"custom_fee"`
}

type Item struct {
	ID          int64   `gorm:"primaryKey;column:item_id"`
	OrderUID    string  `gorm:"column:order_uid"`
	ChrtId      int64   `json:"chrt_id"`
	TrackNumber string  `json:"track_number"`
	Price       float64 `json:"price"`
	Rid         string  `json:"rid"`
	Name        string  `json:"name"`
	Sale        int64   `json:"sale"`
	Size        string  `json:"size" gorm:"column:item_size"`
	TotalPrice  float64 `json:"total_price"`
	NmId        int64   `json:"nm_id"`
	Brand       string  `json:"brand"`
	Status      int64   `json:"status"`
}

// Reads json file and returns Order model
func ReadModel(filename string) (*Order, error) {
	if filename == "" {
		return nil, errors.New("filename can't be empty")
	}
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var order Order
	err = json.Unmarshal(byteValue, &order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func TestOrder(t *testing.T, order *Order) {
	TestField(t, "order_uid", order.UID, "b563feb7b2b84b6test")
	TestField(t, "order_uid", order.UID, "b563feb7b2b84b6test")
	TestField(t, "track_number", order.TrackNumber, "WBILMTESTTRACK")
	TestField(t, "entry", order.Entry, "WBIL")
	TestField(t, "locale", order.Locale, "en")
	TestField(t, "internal_signature", order.InternalSignature, "")
	TestField(t, "customer_id", order.CustomerId, "test")
	TestField(t, "delivery_service", order.DeliveryService, "meest")
	TestField(t, "shardkey", order.Shardkey, "9")
	TestField(t, "sm_id", order.SmId, 99)
	TestField(t, "date_created", order.DateCreated.Format(time.RFC3339), "2021-11-26T06:22:19Z")
	TestField(t, "oof_shard", order.OofShard, "1")

	// Delivery fields
	TestField(t, "delivery.name", order.Delivery.Name, "Test Testov")
	TestField(t, "delivery.phone", order.Delivery.Phone, "+9720000000")
	TestField(t, "delivery.zip", order.Delivery.Zip, "2639809")
	TestField(t, "delivery.city", order.Delivery.City, "Kiryat Mozkin")
	TestField(t, "delivery.address", order.Delivery.Address, "Ploshad Mira 15")
	TestField(t, "delivery.region", order.Delivery.Region, "Kraiot")
	TestField(t, "delivery.email", order.Delivery.Email, "test@gmail.com")

	// Payment fields
	TestField(t, "payment.transaction", order.Payment.Transaction, "b563feb7b2b84b6test")
	TestField(t, "payment.request_id", order.Payment.RequestId, "")
	TestField(t, "payment.currency", order.Payment.Currency, "USD")
	TestField(t, "payment.provider", order.Payment.Provider, "wbpay")
	TestField(t, "payment.amount", order.Payment.Amount, 1817)
	TestField(t, "payment.payment_dt", order.Payment.PaymentDt, 1637907727)
	TestField(t, "payment.bank", order.Payment.Bank, "alpha")
	TestField(t, "payment.delivery_cost", order.Payment.DeliveryCost, 1500)
	TestField(t, "payment.goods_total", order.Payment.GoodsTotal, 317)
	TestField(t, "payment.custom_fee", order.Payment.CustomFee, 0)

	// Items fields (assuming there's only one item in the list)
	item := order.Items[0]
	TestField(t, "items[0].chrt_id", item.ChrtId, 9934930)
	TestField(t, "items[0].track_number", item.TrackNumber, "WBILMTESTTRACK")
	TestField(t, "items[0].price", item.Price, 453)
	TestField(t, "items[0].rid", item.Rid, "ab4219087a764ae0btest")
	TestField(t, "items[0].name", item.Name, "Mascaras")
	TestField(t, "items[0].sale", item.Sale, 30)
	TestField(t, "items[0].size", item.Size, "0")
	TestField(t, "items[0].total_price", item.TotalPrice, 317)
	TestField(t, "items[0].nm_id", item.NmId, 2389212)
	TestField(t, "items[0].brand", item.Brand, "Vivienne Sabo")
	TestField(t, "items[0].status", item.Status, 202)
}

func TestField[T comparable](t *testing.T, name string, actualValue T, expectedValue T) {
	if actualValue != expectedValue {
		t.Errorf("Wrong %s. want=%v, got=%v", name, actualValue, expectedValue)
	}
}
