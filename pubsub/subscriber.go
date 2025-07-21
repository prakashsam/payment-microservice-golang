package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"paymentservice/models"
	"paymentservice/services"

	"cloud.google.com/go/pubsub"
)

func StartOrderSubscriber(ctx context.Context) {
	client, _ := pubsub.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	sub := client.Subscription(os.Getenv("ORDER_SUBSCRIPTION_ID"))
	sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		var order models.Payment
		if err := json.Unmarshal(msg.Data, &order); err != nil {
			msg.Nack()
			return
		}
		services.CreatePayment(order.OrderID, order.Amount)
		msg.Ack()
	})
}
