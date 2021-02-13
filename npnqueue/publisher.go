package npnqueue

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
)

// Publishes messages to a queue
type Publisher struct {
	Topic  string
	Addrs  []string
	Count  int
	Last   time.Time
	writer sarama.SyncProducer
}

// Constructor
func NewPublisher(cfg *Config, logger *logrus.Logger) (*Publisher, error) {
	config := makeSaramaConfig(cfg.Secure, cfg.Username, cfg.Password, cfg.Secure, cfg.Verbose, logger)
	producer, err := sarama.NewSyncProducer(cfg.Addrs, config)
	if err != nil {
		return nil, err
	}

	return &Publisher{Topic: cfg.Topic, Addrs: cfg.Addrs, Count: 0, writer: producer}, nil
}

// Write a message to the queue
func (c *Publisher) Write(ctx context.Context, key string, m *Message) error {
	json := m.Payload
	hd := make([]sarama.RecordHeader, 0, len(m.Headers))
	for k, v := range m.Headers {
		hd = append(hd, sarama.RecordHeader{Key: []byte(k), Value: v})
	}
	message := &sarama.ProducerMessage{
		Topic:     m.Topic,
		Key:       sarama.StringEncoder(m.Key),
		Value:     sarama.StringEncoder(json),
		Headers:   hd,
		Timestamp: time.Now(),
	}
	c.Count += 1
	c.Last = time.Now()
	_, _, err := c.writer.SendMessage(message)
	return err
}

func (c *Publisher) Close() error {
	return nil
}
