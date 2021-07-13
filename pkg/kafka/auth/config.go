package auth

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"go.uber.org/zap"

)

const (
	none ="none"
	kerberos = "kerberos"
	tls = "tls"
	plaintext = "plaintext"
)

var authTypes = []string {
	none,
	kerberos,
	tls,
	plaintext
}

type AuthenticationConfig struct {
	Authentication string
	Kerberos 
}