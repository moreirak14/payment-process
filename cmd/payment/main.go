package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/moreirak14/payment-process/internal/entity"
	"github.com/moreirak14/payment-process/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log/slog"
)

func main() {
	ctx := context.Background()
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	messages := make(chan amqp.Delivery)
	// Consume from the queue "payments" and send the messages to the channel "messages"
	// This is a blocking call
	// Go to pkg/rabbitmq/rabbitmq.go to see the implementation
	go rabbitmq.Consume(channel, "payments", messages) // Go routine

	for message := range messages {
		// Unmarshal the message
		var orderRequest entity.OrderPaymentRequest

		err := json.Unmarshal(message.Body, &orderRequest)
		if err != nil {
			slog.Error(err.Error())
			break
		}

		// Process the order
		response, err := orderRequest.Process()
		if err != nil {
			slog.Error(err.Error())
		}
		fmt.Println(response)

		// Marshal the response
		responseJSON, err := json.Marshal(response)
		if err != nil {
			slog.Error(err.Error())
			break
		}

		// Publish the response to the exchange "amqp.direct" with the routing key "payments
		err = rabbitmq.Publish(ctx, channel, string(responseJSON), "amqp.direct")
		if err != nil {
			slog.Error(err.Error())
			break
		}

		// Acknowledge the message
		message.Ack(false)

		slog.Info("Message processed")
	}

	defer channel.Close()
}

//json = {
//"viagemExternoId": 1279,
//"tipo": "Complemento",
//"formaPagamento": "Pix",
//"cpfCnpjContratado": "83250794872",
//"nomeContratado": "Kaique Moreira",
//"chavePix": "39606753808",
//"cpfMotorista": "39606753808",
//"nomeMotorista": "Kaique Moreira",
//"valor": 5.0,
//"tipoBanco": "BBC",
//"pagamentoExternoId": 3955,
//"ibgeOrigem": 3538709,
//"ibgeDestino": 3538709,
//"hashValidacao": "37c30035d18976f5590c2a525e02cf4c",
//}
