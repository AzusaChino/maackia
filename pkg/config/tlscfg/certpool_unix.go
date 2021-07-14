package tlscfg

import "crypto/x509"

func loadSystemCertPool() (*x509.CertPool, error) {
	return systemCertPool()
}