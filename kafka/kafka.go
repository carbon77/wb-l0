package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"ru/zakat/L0/cache"
	"ru/zakat/L0/logger"
	"ru/zakat/L0/models"

	"github.com/IBM/sarama"
)

var (
	log = logger.NewLogger()
)

type consumerGroupHandler struct {
	cache *cache.Cache
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Info("Received message from \"orders\" topic")
		var order models.Order
		err := json.Unmarshal(msg.Value, &order)

		if err != nil {
			log.Error("Can't read order from message")
		} else {
			log.Info("Order has been read")
			h.cache.AddOrder(&order)
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func NewConsumer(cache *cache.Cache) {
	group, err := sarama.NewConsumerGroup([]string{"localhost:9094"}, "my-group", nil)
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
