package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/solution-labs/tools/toolserror"
)

type PBMessage struct {
	ProjectID   string
	Topic       string
	Attribs     map[string]string
	Message     interface{}
	OrderingKey string
}

func SendMessage(ctx context.Context, pb PBMessage) (pr string, err error) {

	client, err := pubsub.NewClient(ctx, pb.ProjectID)
	if err != nil {
		return pr, toolserror.Wrap("pubsub:SendMessage: %w", err)
	}

	topic := client.Topic(pb.Topic)

	dataJson, err := json.Marshal(pb.Message)
	if err != nil {
		return pr, toolserror.Wrap("pubsub:SendMessage:Json: %w", err)
	}

	msg := &pubsub.Message{
		Data:        dataJson,
		Attributes:  pb.Attribs,
		OrderingKey: pb.OrderingKey,
	}

	if pr, err = topic.Publish(ctx, msg).Get(ctx); err != nil {
		return pr, toolserror.Wrap("pubsub:SendMessage:Publish: %w", err)
	}

	return pr, nil
}
