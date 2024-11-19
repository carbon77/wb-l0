package models

import (
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
