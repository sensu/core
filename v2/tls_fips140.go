//go:build fips140
// +build fips140

package v2

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

var (
	// Allow TLS 1.2 and TLS 1.3
	tlsMinVersion = uint16(tls.VersionTLS12)
	tlsMaxVersion = uint16(tls.VersionTLS13)

	// FIPS-approved TLS 1.2 cipher suites
	DefaultCipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	}

	// FIPS-approved NIST curves
	tlsCurvePreferences = []tls.CurveID{
		tls.CurveP384,
		tls.CurveP256,
	}
)

// ToServerTLSConfig creates a hardened TLS configuration for servers
func (t *TLSOptions) ToServerTLSConfig() (*tls.Config, error) {
	cfg := &tls.Config{
		MinVersion:               tlsMinVersion,
		MaxVersion:               tlsMaxVersion,
		CipherSuites:             DefaultCipherSuites,
		CurvePreferences:         tlsCurvePreferences,
		PreferServerCipherSuites: true,
		SessionTicketsDisabled:   true,
		Renegotiation:            tls.RenegotiateNever,
	}

	// Load CA certificates for client authentication
	if t.GetTrustedCAFile() != "" {
		caCertPool, err := LoadCACerts(t.TrustedCAFile)
		if err != nil {
			return nil, err
		}
		cfg.ClientCAs = caCertPool
	}

	// Load server certificate
	if t.GetCertFile() != "" && t.GetKeyFile() != "" {
		cert, err := tls.LoadX509KeyPair(t.GetCertFile(), t.GetKeyFile())
		if err != nil {
			return nil, fmt.Errorf("error loading TLS server certificate: %w", err)
		}

		cfg.Certificates = []tls.Certificate{cert}
	}

	// Enable mutual TLS if configured
	if t.GetClientAuthType() {
		cfg.ClientAuth = tls.RequireAndVerifyClientCert
	}

	return cfg, nil
}

// ToClientTLSConfig creates a hardened TLS configuration for clients
func (t *TLSOptions) ToClientTLSConfig() (*tls.Config, error) {
	cfg := &tls.Config{
		MinVersion:             tlsMinVersion,
		MaxVersion:             tlsMaxVersion,
		CipherSuites:           DefaultCipherSuites,
		CurvePreferences:       tlsCurvePreferences,
		SessionTicketsDisabled: true,
		Renegotiation:          tls.RenegotiateNever,
		InsecureSkipVerify:     false,
	}

	// Load trusted CA certificates
	if t.GetTrustedCAFile() != "" {
		caCertPool, err := LoadCACerts(t.TrustedCAFile)
		if err != nil {
			return nil, err
		}
		cfg.RootCAs = caCertPool
	}

	// Load client certificate
	if t.GetCertFile() != "" && t.GetKeyFile() != "" {
		cert, err := tls.LoadX509KeyPair(t.GetCertFile(), t.GetKeyFile())
		if err != nil {
			return nil, fmt.Errorf("error loading TLS client certificate: %w", err)
		}

		cfg.Certificates = []tls.Certificate{cert}
	}

	return cfg, nil
}

// LoadCACerts loads CA certificates from a PEM bundle
func LoadCACerts(path string) (*x509.CertPool, error) {
	caCerts, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading CA file: %w", err)
	}

	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caCerts) {
		return nil, fmt.Errorf("no certificates could be parsed from %s", path)
	}

	return caCertPool, nil
}