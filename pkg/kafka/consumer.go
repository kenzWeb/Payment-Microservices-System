package kafka

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Consumer struct {
	client *kgo.Client
}

func NewConsumer(brokers []string, groupID string, topics []string) (*Consumer, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ConsumerGroup(groupID),
		kgo.ConsumeTopics(topics...),
	)
	if err != nil {
		return nil, fmt.Errorf("new kafka client: %w", err)
	}
	return &Consumer{client: cl}, nil
}

func (c *Consumer) Poll(ctx context.Context) ([]*kgo.Record, error) {
	fetches := c.client.PollRecords(ctx, 1)
	if err := fetches.Err(); err != nil {
		return nil, err
	}
	return fetches.Records(), nil
}

func (c *Consumer) Close() {
	c.client.Close()
}
