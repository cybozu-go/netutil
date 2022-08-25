package netutil

import (
	"net"
	"testing"
)

func TestIP4ToInt(t *testing.T) {
	t.Parallel()

	if IP4ToInt(net.ParseIP("0.0.1.0")) != 256 {
		t.Error(`0.0.1.0 != 256`)
	}

	if IP4ToInt(net.ParseIP("0.0.1.1")) != 257 {
		t.Error(`0.0.1.0 != 257`)
	}
}

func TestIntToIP4(t *testing.T) {
	t.Parallel()

	if !IntToIP4(256).Equal(net.ParseIP("0.0.1.0")) {
		t.Error(`256 != 0.0.1.0`)
	}

	if !IntToIP4(257).Equal(net.ParseIP("0.0.1.1")) {
		t.Error(`257 != 0.0.1.0`)
	}
}

func TestHostsFunc(t *testing.T) {
	t.Parallel()

	_, n, err := net.ParseCIDR("12.34.56.0/30")
	if err != nil {
		t.Fatal(err)
	}

	f, err := HostsFunc(n)
	if err != nil {
		t.Fatal(err)
	}

	if !f().Equal(net.ParseIP("12.34.56.1")) {
		t.Error("f() != 12.34.56.1")
	}

	if !f().Equal(net.ParseIP("12.34.56.2")) {
		t.Error("f() != 12.34.56.2")
	}

	if f() != nil {
		t.Error("f() != nil")
	}
}
