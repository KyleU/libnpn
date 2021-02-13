package npnqueue

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"github.com/kyleu/libnpn/npncore"
	"github.com/sirupsen/logrus"

	"github.com/Shopify/sarama"
)

// A queue message
type Message struct {
	Topic   string            `json:"topic"`
	Key     string            `json:"key"`
	Headers map[string][]byte `json:"headers,omitempty"`
	Payload json.RawMessage   `json:"payload"`
	Time    time.Time         `json:"time,omitempty"`
}

// Array helper
type Messages []*Message

// Queue configuration, loaded from environment
type Config struct {
	Secure   bool
	Addrs    []string
	Username string
	Password string
	Topic    string
	Verbose  bool
}

// Loads the queue configuration from the environment
func LoadConfig(verbose bool, defaultTopic string) *Config {
	secure := env("secure", "true") != "false"
	addrs := strings.Split(env("host", "localhost:9092"), ",")
	u := env("username", "")
	p := env("password", "")
	topic := env("topic", defaultTopic)
	return &Config{Secure: secure, Addrs: addrs, Username: u, Password: p, Topic: topic, Verbose: verbose}
}

func env(key string, dflt string) string {
	x, ok := os.LookupEnv(npncore.AppKey + "_kafka_" + key)
	if ok && len(x) > 0 {
		return x
	}
	return dflt
}

func makeSaramaConfig(secure bool, username string, password string, sasl bool, verbose bool, logger *logrus.Logger) *sarama.Config {
	if verbose {
		sarama.Logger = logger
	}

	config := sarama.NewConfig()
	config.ClientID = npncore.AppKey
	if len(username) > 0 && username != "null" {
		config.Net.SASL.Enable = sasl
		config.Net.SASL.User = username
		config.Net.SASL.Password = password
	}
	config.Net.TLS.Enable = secure
	config.Producer.Return.Successes = true
	return config
}
