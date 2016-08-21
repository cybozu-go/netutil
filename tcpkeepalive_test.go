package netutil

import (
	"bytes"
	"io/ioutil"
	"net"
	"os"
	"testing"
	"time"
)

const (
	testAddress = "localhost:15345"
)

func TestTCPKeepAlive(t *testing.T) {
	t.Parallel()

	addr, err := net.ResolveTCPAddr("tcp", testAddress)
	if err != nil {
		t.Fatal(err)
	}

	tl, err := net.ListenTCP("tcp", addr)
	if err != nil {
		t.Skip(err)
	}

	go func(l net.Listener) {
		c, err := l.Accept()
		l.Close()
		if err != nil {
			t.Error(err)
			return
		}
		c.Write([]byte("abc"))
		c.Close()
	}(TCPKeepAliveListener{tl})

	c, err := net.DialTimeout("tcp", testAddress, 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	SetKeepAlive(c)

	data, err := ioutil.ReadAll(c)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, []byte("abc")) {
		t.Error(`!bytes.Equal(data, []byte("abc"))`)
	}
}

func TestKeepAliveListener(t *testing.T) {
	t.Parallel()

	f, err := ioutil.TempFile("", "gotest")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	os.Remove(f.Name())

	l, err := net.Listen("unix", f.Name())
	if err != nil {
		t.Fatal(err)
	}

	if l != KeepAliveListener(l) {
		t.Error(`l != KeepAliveListener(l)`)
	}

	go func() {
		c, err := l.Accept()
		l.Close()
		if err != nil {
			t.Error(err)
		}
		c.Write([]byte("xyz"))
		c.Close()
	}()

	c, err := net.DialTimeout("unix", f.Name(), 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	SetKeepAlive(c) // noop for UNIX-domain socket

	data, err := ioutil.ReadAll(c)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, []byte("xyz")) {
		t.Error(`!bytes.Equal(data, []byte("xyz"))`)
	}
}
