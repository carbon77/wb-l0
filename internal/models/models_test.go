package models

import (
	"testing"
	"time"
)

func TestEmptyFilename(t *testing.T) {
	order, err := ReadModel("")
	if err == nil {
		t.Fatal("Error is empty")
	}

	if order != nil {
		t.Fatal("Order is not empty")
	}

	if err.Error() != "filename can't be empty" {
		t.Fatalf("wrong error")
	}
}

// Testing ReadModel function
func TestReadModel(t *testing.T) {
	filename := "../../model.json"
	order, err := ReadModel(filename)
	if err != nil {
		t.Fatal(err)
	}

	testField(t, "order_uid", order.UID, "b563feb7b2b84b6test")
	testField(t, "order_uid", order.UID, "b563feb7b2b84b6test")
	testField(t, "track_number", order.TrackNumber, "WBILMTESTTRACK")
	testField(t, "entry", order.Entry, "WBIL")
	testField(t, "locale", order.Locale, "en")
	testField(t, "internal_signature", order.InternalSignature, "")
	testField(t, "customer_id", order.CustomerId, "test")
	testField(t, "delivery_service", order.DeliveryService, "meest")
	testField(t, "shardkey", order.Shardkey, "9")
	testField(t, "sm_id", order.SmId, 99)
	testField(t, "date_created", order.DateCreated.Format(time.RFC3339), "2021-11-26T06:22:19Z")
	testField(t, "oof_shard", order.OofShard, "1")

	// Delivery fields
	testField(t, "delivery.name", order.Delivery.Name, "Test Testov")
	testField(t, "delivery.phone", order.Delivery.Phone, "+9720000000")
	testField(t, "delivery.zip", order.Delivery.Zip, "2639809")
	testField(t, "delivery.city", order.Delivery.City, "Kiryat Mozkin")
	testField(t, "delivery.address", order.Delivery.Address, "Ploshad Mira 15")
	testField(t, "delivery.region", order.Delivery.Region, "Kraiot")
	testField(t, "delivery.email", order.Delivery.Email, "test@gmail.com")

	// Payment fields
	testField(t, "payment.transaction", order.Payment.Transaction, "b563feb7b2b84b6test")
	testField(t, "payment.request_id", order.Payment.RequestId, "")
	testField(t, "payment.currency", order.Payment.Currency, "USD")
	testField(t, "payment.provider", order.Payment.Provider, "wbpay")
	testField(t, "payment.amount", order.Payment.Amount, 1817)
	testField(t, "payment.payment_dt", order.Payment.PaymentDt, 1637907727)
	testField(t, "payment.bank", order.Payment.Bank, "alpha")
	testField(t, "payment.delivery_cost", order.Payment.DeliveryCost, 1500)
	testField(t, "payment.goods_total", order.Payment.GoodsTotal, 317)
	testField(t, "payment.custom_fee", order.Payment.CustomFee, 0)

	// Items fields (assuming there's only one item in the list)
	item := order.Items[0]
	testField(t, "items[0].chrt_id", item.ChrtId, 9934930)
	testField(t, "items[0].track_number", item.TrackNumber, "WBILMTESTTRACK")
	testField(t, "items[0].price", item.Price, 453)
	testField(t, "items[0].rid", item.Rid, "ab4219087a764ae0btest")
	testField(t, "items[0].name", item.Name, "Mascaras")
	testField(t, "items[0].sale", item.Sale, 30)
	testField(t, "items[0].size", item.Size, "0")
	testField(t, "items[0].total_price", item.TotalPrice, 317)
	testField(t, "items[0].nm_id", item.NmId, 2389212)
	testField(t, "items[0].brand", item.Brand, "Vivienne Sabo")
	testField(t, "items[0].status", item.Status, 202)

}

func testField[T comparable](t *testing.T, name string, actualValue T, expectedValue T) {
	if actualValue != expectedValue {
		t.Errorf("Wrong %s. want=%v, got=%v", name, actualValue, expectedValue)
	}
}
