package netutil

import (
	"errors"
	"net"
	"syscall"
	"testing"
)

func TestErrors(t *testing.T) {
	t.Parallel()

	if !IsNetworkUnreachable(syscall.ENETUNREACH) {
		t.Error(`!IsNetworkUnreachable(syscall.ENETUNREACH)`)
	}
	if IsNetworkUnreachable(errors.New("hoge")) {
		t.Error(`IsNetworkUnreachable(errors.New("hoge"))`)
	}

	if !IsConnectionRefused(syscall.ECONNREFUSED) {
		t.Error(`!IsConnectionRefused(syscall.ECONNREFUSED)`)
	}
	if IsConnectionRefused(errors.New("hoge")) {
		t.Error(`IsConnectionRefused(errors.New("hoge"))`)
	}

	if !IsNoRouteToHost(syscall.EHOSTUNREACH) {
		t.Error(`!IsNoRouteToHost(syscall.EHOSTUNREACH)`)
	}
	if IsNoRouteToHost(errors.New("hoge")) {
		t.Error(`IsNoRouteToHost(errors.New("hoge"))`)
	}
}

func TestConnectionRefused(t *testing.T) {
	t.Parallel()

	_, err := net.Dial("tcp", "localhost:38794")
	if err == nil {
		return
	}
	if !IsConnectionRefused(err) {
		t.Error(`!IsConnectionRefused(err)`)
		t.Log(err.Error())
	}
}
