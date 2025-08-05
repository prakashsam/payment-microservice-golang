package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"paymentservice/config"
	"paymentservice/models"
	"paymentservice/services"

	"cloud.google.com/go/pubsub"
)

func StartOrderSubscriber(ctx context.Context, paymentService *services.PaymentService) {
	cfg := config.Load()
	client, _ := pubsub.NewClient(ctx, cfg.ProjectID)
	sub := client.Subscription(cfg.orderSubscriptionID)
	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		var order models.Payment
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			msg.Nack()
			return
		}
		paymentService.CreatePayment(order.OrderID, order.Amount)
		msg.Ack()
	})
}

