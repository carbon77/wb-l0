package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"ru/zakat/L0/internal/cache"
	"ru/zakat/L0/internal/config"
	"ru/zakat/L0/internal/logger"
	"ru/zakat/L0/internal/models"

	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

var (
	log     = logger.NewLogger()
	address = []string{config.KafkaUrl}
)

type OrderProducer struct {
	producer sarama.SyncProducer
}

type consumerGroupHandler struct {
	cache *cache.Cache
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Info("Received message", zap.String("topic", "orders"))
		var order models.Order
		err := json.Unmarshal(msg.Value, &order)

		if err != nil {
			log.Error("Can't read order from message", zap.Error(err))
		} else {
			log.Info("Order has been read", zap.String("order_uid", order.UID))
			h.cache.AddOrder(&order)
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func NewConsumer(cache *cache.Cache) {
	group, err := sarama.NewConsumerGroup(address, "my-group", nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = group.Close() }()

	go func() {
		for err := range group.Errors() {
			fmt.Println("ERROR", err)
		}
	}()

	ctx := context.Background()
	for {
		topics := []string{"orders"}
		handler := consumerGroupHandler{cache}

		err := group.Consume(ctx, topics, handler)
		if err != nil {
			panic(err)
		}
	}
}

func NewProducer() *OrderProducer {
	producer, err := sarama.NewSyncProducer(address, nil)
	if err != nil {
		log.Error("Can't create sync producer", zap.Error(err))
	}

	return &OrderProducer{
		producer: producer,
	}
}

func (p *OrderProducer) SendOrder(topic string, order *models.Order) {
	bytes, err := json.Marshal(order)
	if err != nil {
		log.Error("Can't marshall order", zap.Error(err))
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(order.UID),
		Value: sarama.ByteEncoder(bytes),
	}

	_, _, err = p.producer.SendMessage(msg)
	if err != nil {
		log.Error("Failed to send message", zap.Error(err))
	} else {
		log.Info("Order sent", zap.String("order_uid", order.UID))
	}
}
