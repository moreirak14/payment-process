package gcp

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func ClientPubSub(projectID string) (*pubsub.Client, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}
	return client, nil
}

func PublishMessage(client *pubsub.Client, topicID, msg string) error {
	ctx := context.Background()

	topic := client.Topic(topicID)
	exist, err := topic.Exists(ctx)
	if err != nil {
		return fmt.Errorf("topic.Exists: %v", err)

	}
	if !exist {
		topic, err = client.CreateTopic(ctx, topicID)
		if err != nil {
			return fmt.Errorf("CreateTopic: %v", err)
		}
	}

	publish := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// The publish happens asynchronously.
	// Later, you can get the result from result.Get.
	id, err := publish.Get(ctx)
	if err != nil {
		return fmt.Errorf("Get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

func ConsumeMessage(client *pubsub.Client, subscriptionID string) error {
	ctx := context.Background()

	sub := client.Subscription(subscriptionID)
	_, err = sub.Exists(ctx)
	if err != nil {
		return fmt.Errorf("sub.Exists: %v", err)

	}

	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("Receive: %v", err)
	}
	return nil
}
