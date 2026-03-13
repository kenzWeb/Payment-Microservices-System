package kafka

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
)

type Producer struct {
	client *kgo.Client
}

func NewProducer(brokers []string) (*Producer, error) {
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.AllowAutoTopicCreation(),
		kgo.RequiredAcks(kgo.AllISRAcks()),
		kgo.RecordRetries(5),
	)
	if err != nil {
		return nil, fmt.Errorf("kafka client: %w", err)
	}
	return &Producer{client: cl}, nil
}

func (p *Producer) Publish(ctx context.Context, topic string, key, value []byte) error {
	record := &kgo.Record{Topic: topic, Key: key, Value: value}
	return p.client.ProduceSync(ctx, record).FirstErr()
}

func (p *Producer) Close() {
	p.client.Close()
}
