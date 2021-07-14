package consumer

import (
	"io"
	"time"

	"github.com/AzusaChino/makia/pkg/kafka/auth"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"go.uber.org/zap"
)

type Consumer interface {
	Partitions() <-chan cluster.PartitionConsumer
	MarkPartitionOffset(topic string, partition int32, offset int64, metadata string)
	io.Closer
}

type Builder interface {
	NewConsumer() (Consumer, error)
}

// Configuration describes the configuration properties needed to create a Kafka consumer
type Configuration struct {
	auth.AuthenticationConfig `mapstructure:"authentication"`
	Consumer

	Brokers         []string `mapstructure:"brokers"`
	Topic           string   `mapstructure:"topic"`
	GroupId         string   `mapstructure:"group_id"`
	ClientId        string   `mapstructure:"client_id"`
	ProtocolVersion string   `mapstructure:"protocol_version"`
}

func (c *Configuration) NewConsumer(logger *zap.Logger) (Consumer, error) {
	saramaConfig := cluster.NewConfig()
	saramaConfig.Group.Mode = cluster.ConsumerModePartitions
	saramaConfig.ClientID = c.ClientId
	if len(c.ProtocolVersion) > 0 {
		version, err := sarama.ParseKafkaVersion(c.ProtocolVersion)
		if err != nil {
			return nil, err
		}
		saramaConfig.Config.Version = version
	}
	if err := c.AuthenticationConfig.SetConfiguration(&saramaConfig.Config, logger); err != nil {
		return nil, err
	}
	saramaConfig.Consumer.Offsets.AutoCommit.Enable = true
	saramaConfig.Consumer.Offsets.AutoCommit.Interval = time.Second
	return cluster.NewConsumer(c.Brokers, c.GroupId, []string{c.Topic}, saramaConfig)
}
