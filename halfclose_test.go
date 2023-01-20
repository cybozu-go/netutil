package netutil

import (
	"bytes"
	"fmt"
	"io"
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

	errCh := make(chan error, 1)
	done := make(chan struct{})

	go func() {
		defer close(done)
		c, err := l.Accept()
		if err != nil {
			errCh <- err
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
			errCh <- err
			return
		}
		data, err := io.ReadAll(c)
		if err != nil {
			errCh <- err
		} else if !bytes.Equal(data, []byte("ack")) {
			errCh <- fmt.Errorf(`!bytes.Equal(data, []byte("ack")), data=%s`, data)
		}
	}()

	c, err := net.DialTimeout("tcp", "localhost:15346", 1*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	data, err := io.ReadAll(c)
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
	select {
	case err2 := <-errCh:
		t.Error(err2)
	default:
	}
}
