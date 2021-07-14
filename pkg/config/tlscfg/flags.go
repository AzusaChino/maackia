package tlscfg

import (
	"flag"
	"github.com/spf13/viper"
)

const (
	tlsPrefix         = ".tls"
	tlsEnabled        = tlsPrefix + ".enabled"
	tlsCA             = tlsPrefix + ".ca"
	tlsCert           = tlsPrefix + ".cert"
	tlsKey            = tlsPrefix + ".key"
	tlsServerName     = tlsPrefix + ".server-name"
	tlsClientCA       = tlsPrefix + ".client-ca"
	tlsSkipHostVerify = tlsPrefix + ".skip-host-verify"
)

// tls client config for cli flags
type ClientFlagsConfig struct {
	Prefix string
}

// tls server config for cli flags
type ServerFlagsConfig struct {
	Prefix string
}

// AddFlags adds flags for TLS to the FlagSet.
func (c ClientFlagsConfig) AddFlags(flags *flag.FlagSet) {
	flags.Bool(c.Prefix+tlsEnabled, false, "Enable TLS when talking to the remote server(s)")
	flags.String(c.Prefix+tlsCA, "", "Path to a TLS CA (Certification Authority) file used to verify the remote server(s) (by default will use the system truststore)")
	flags.String(c.Prefix+tlsCert, "", "Path to a TLS Certificate file, used to identify this process to the remote server(s)")
	flags.String(c.Prefix+tlsKey, "", "Path to a TLS Private Key file, used to identify this process to the remote server(s)")
	flags.String(c.Prefix+tlsServerName, "", "Override the TLS server name we expect in the certificate of the remote server(s)")
	flags.Bool(c.Prefix+tlsSkipHostVerify, false, "(insecure) Skip server's certificate chain and host name verification")
}

// AddFlags adds flags for TLS to the FlagSet.
func (c ServerFlagsConfig) AddFlags(flags *flag.FlagSet) {
	flags.Bool(c.Prefix+tlsEnabled, false, "Enable TLS on the server")
	flags.String(c.Prefix+tlsCert, "", "Path to a TLS Certificate file, used to identify this server to clients")
	flags.String(c.Prefix+tlsKey, "", "Path to a TLS Private Key file, used to identify this server to clients")
	flags.String(c.Prefix+tlsClientCA, "", "Path to a TLS CA (Certification Authority) file used to verify certificates presented by clients (if unset, all clients are permitted)")
}

// InitFromViper creates tls.Config populated with values retrieved from Viper.
func (c ClientFlagsConfig) InitFromViper(v *viper.Viper) Options {
	var p Options
	p.Enabled = v.GetBool(c.Prefix + tlsEnabled)
	p.CAPath = v.GetString(c.Prefix + tlsCA)
	p.CertPath = v.GetString(c.Prefix + tlsCert)
	p.KeyPath = v.GetString(c.Prefix + tlsKey)
	p.ServerName = v.GetString(c.Prefix + tlsServerName)
	p.SkipHostVerify = v.GetBool(c.Prefix + tlsSkipHostVerify)
	return p
}

// InitFromViper creates tls.Config populated with values retrieved from Viper.
func (c ServerFlagsConfig) InitFromViper(v *viper.Viper) Options {
	var p Options
	p.Enabled = v.GetBool(c.Prefix + tlsEnabled)
	p.CertPath = v.GetString(c.Prefix + tlsCert)
	p.KeyPath = v.GetString(c.Prefix + tlsKey)
	p.ClientCAPath = v.GetString(c.Prefix + tlsClientCA)
	return p
}
