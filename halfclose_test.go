package netutil

import (
	"bytes"
	"io/ioutil"
	"net"
	"testing"
	"time"
)

func TestHalfCloser(t *testing.T) {
	t.Parallel()

	l, err := net.Listen("tcp", "localhost:15346")
	if err != nil {
		t.Skip(err)
	}

	done := make(chan struct{})

	go func() {
		c, err := l.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer c.Close()

		hc, ok := c.(HalfCloser)
		if !ok {
			t.Error("c is not a HalfCloser")
		}
		c.Write([]byte("test"))
		err = hc.CloseWrite()
		if err != nil {
			t.Fatal(err)
		}
		data, err := ioutil.ReadAll(c)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(data, []byte("ack")) {
			t.Error(`!bytes.Equal(data, []byte("ack"))`)
		}
		close(done)
	}()

	c, err := net.DialTimeout("tcp", "localhost:15346", 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	data, err := ioutil.ReadAll(c)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, []byte("test")) {
		t.Error(`!bytes.Equal(data, []byte("test"))`)
	}
	_, err = c.Write([]byte("ack"))
	if err != nil {
		t.Error(err)
	}
	err = c.(HalfCloser).CloseWrite()
	if err != nil {
		t.Error(err)
	}

	<-done
}
