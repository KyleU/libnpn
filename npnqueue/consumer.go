package npnqueue

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"logur.dev/logur"
)

// A queue consumer, for reading messages
type Consumer struct {
	Topic   string `json:"topic,omitempty"`
	Addrs   []string
	Reader  sarama.ConsumerGroup
	Group   string
	Handler *ConsumeHelper
	Config  *sarama.Config
	Count   int
	Last    time.Time
	logger  logur.Logger
}

// Creates a new Consumer from the provided ConsumeHelper
func NewConsumer(cfg *Config, group string, handler *ConsumeHelper, logger logur.Logger, logr *logrus.Logger) (*Consumer, error) {
	config := makeSaramaConfig(cfg.Secure, cfg.Username, cfg.Password, cfg.Secure, cfg.Verbose, logr)
	r, err := sarama.NewConsumerGroup(cfg.Addrs, group, config)
	if err != nil {
		return nil, err
	}
	return &Consumer{Topic: cfg.Topic, Addrs: cfg.Addrs, Reader: r, Group: group, Handler: handler, Config: config, logger: logger}, nil
}

// Enters a loop that reads messages until the connection is closed
func (c *Consumer) ReadLoop(onMessage func(msg *Message), onError func(e error)) {
	go func() {
		for {
			err := c.Read(context.Background())
			if err != nil {
				c.logger.Error(fmt.Sprintf("error reading: %+v", err))
			}
			c.Reset()
		}
	}()

	for {
		select {
		case msg := <-c.Handler.Consumers:
			c.Count += 1
			c.Last = time.Now()
			go onMessage(msg)
		case err := <-c.Handler.Errors:
			go onError(err)
		}
	}
}

// Reads a single message from the queue
func (c *Consumer) Read(ctx context.Context) error {
	return c.consume(ctx)
}

// Closes the queue connection
func (c *Consumer) Close() error {
	// return c.Reader.Close()
	return nil
}

// Resets the queue connection
func (c *Consumer) Reset() {
	c.Handler.Ready = make(chan bool)
}

func (c *Consumer) consume(ctx context.Context) error {
	return c.Reader.Consume(ctx, []string{c.Topic}, c.Handler)
}

// Channels used by the Consumer
type ConsumeHelper struct {
	Consumers chan *Message
	Errors    chan error
	Ready     chan bool
}

// Constructor
func NewConsumeHelper() *ConsumeHelper {
	return &ConsumeHelper{Consumers: make(chan *Message), Errors: make(chan error), Ready: make(chan bool)}
}

// Performs initial setup
func (c *ConsumeHelper) Setup(sarama.ConsumerGroupSession) error {
	close(c.Ready)
	return nil
}

// Unused
func (c *ConsumeHelper) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// Consumes messages from a provided claim
func (c *ConsumeHelper) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for m := range claim.Messages() {
		hd := make(map[string][]byte, len(m.Headers))
		for _, v := range m.Headers {
			hd[string(v.Key)] = v.Value
		}
		msg := &Message{
			Topic:   m.Topic,
			Key:     string(m.Key),
			Headers: hd,
			Payload: m.Value,
			Time:    m.Timestamp,
		}
		c.Consumers <- msg
		session.MarkMessage(m, "")
	}

	return nil
}
