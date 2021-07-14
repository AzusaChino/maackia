package auth

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/xdg-go/scram"
)

// scramClient is the client to use when the auth mechanism is SCRAM
type scramClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

func (c *scramClient) Begin(userName, password, authId string) (err error) {
	c.Client, err = c.HashGeneratorFcn.NewClient(userName, password, authId)
	if err != nil {
		return
	}
	c.ClientConversation = c.Client.NewConversation()
	return nil
}

func (c *scramClient) Step(challenge string) (res string, err error) {
	res, err = c.ClientConversation.Step(challenge)
	return
}

func (c *scramClient) Done() bool {
	return c.ClientConversation.Done()
}

// PlainTextConfig describes the configuration properties needed for SASL/PLAIN with kafka
type PlainTextConfig struct {
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password" json:"-"`
	Mechanism string `mapstructure:"mechanism"`
}

var _ sarama.SCRAMClient = (*scramClient)(nil)

func setPlainTextConfiguration(config *PlainTextConfig, saramaConfig *sarama.Config) error {
	saramaConfig.Net.SASL.Enable = true
	saramaConfig.Net.SASL.User = config.Username
	saramaConfig.Net.SASL.Password = config.Password
	switch strings.ToUpper(config.Mechanism) {
	case "SCRAM-SHA-256":
		saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &scramClient{HashGeneratorFcn: func() hash.Hash { return sha256.New() }}
		}
		saramaConfig.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA256
	case "SCRAM-SHA-512":
		saramaConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &scramClient{HashGeneratorFcn: func() hash.Hash { return sha512.New() }}
		}
		saramaConfig.Net.SASL.Mechanism = sarama.SASLTypeSCRAMSHA512
	case "PLAIN":
		saramaConfig.Net.SASL.Mechanism = sarama.SASLTypePlaintext

	default:
		return fmt.Errorf("config plaintext.mechanism error: %s, only support 'SCRAM-SHA-256' or 'SCRAM-SHA-512' or 'PLAIN'", config.Mechanism)

	}
	return nil
}
