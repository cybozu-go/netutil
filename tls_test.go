package netutil

import (
	"crypto/tls"
	"testing"
)

func TestCipherSuiteString(t *testing.T) {
	t.Parallel()

	if CipherSuiteString(tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA) != "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA" {
		t.Error(`CipherSuiteString(tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA) != "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"`)
	}
}

func TestTLSVersionString(t *testing.T) {
	t.Parallel()

	if TLSVersionString(tls.VersionTLS12) != "TLS1.2" {
		t.Error(`TLSVersionString(tls.VersionTLS12) != "TLS1.2"`)
	}
}
