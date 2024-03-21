package main

import (
	"cloud.google.com/go/pubsub"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/moreirak14/payment-process/pkg/gcp"
	"log/slog"
	"math/rand"
)

type Payment struct {
	PaymentId string `json:"paymentId"`
	Amount    int    `json:"amount"`
}

func main() {
	client, err := connectToPubSub()
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer client.Close()

	msg, err := createPaymentMessage()
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = publishMessage(client, msg)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	err = consumeMessage(client)
	if err != nil {
		slog.Error(err.Error())
	}
}

func connectToPubSub() (*pubsub.Client, error) {
	return gcp.ClientPubSub("projectID")
}

func createPaymentMessage() (string, error) {
	paymentUuid := uuid.New().String()
	randomNumber := rand.New(rand.NewSource(10000))
	payment := Payment{
		PaymentId: paymentUuid,
		Amount:    randomNumber.Int(),
	}

	jsonMsg, err := json.Marshal(payment)
	if err != nil {
		return "", err
	}

	return string(jsonMsg), nil
}

func publishMessage(client *pubsub.Client, msg string) error {
	return gcp.PublishMessage(client, "topicID", msg)
}

func consumeMessage(client *pubsub.Client) error {
	return gcp.ConsumeMessage(client, "subscriptionID")
}
