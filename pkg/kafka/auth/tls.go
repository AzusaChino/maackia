package auth

import (
	"fmt"

	"github.com/AzusaChino/makia/pkg/config/tlscfg"
	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

// set tls config to samara
func setTLSConfiguration(config *tlscfg.Options, samaraConfig *sarama.Config, logger *zap.Logger) error {
	if config.Enabled {
		tlsConfig, err := config.Config(logger)
		if err != nil {
			return fmt.Errorf("error loading tls config: %w", err)
		}
		samaraConfig.Net.TLS.Enable = true
		samaraConfig.Net.TLS.Config = tlsConfig
	}
	return nil
}
