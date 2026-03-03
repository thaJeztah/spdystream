package spdy

import (
	"net"
	"testing"
)

func TestNewFramerWithOptions(t *testing.T) {
	serverConn, clientConn := net.Pipe()
	defer func() {
		_ = clientConn.Close()
		_ = serverConn.Close()
	}()

	framer, err := NewFramerWithOptions(clientConn, clientConn,
		WithMaxControlFramePayloadSize(1024),
		WithMaxHeaderFieldSize(128),
		WithMaxHeaderCount(16),
	)
	if err != nil {
		t.Fatalf("Error creating spdy connection with options: %s", err)
	}

	if got := framer.maxFrameLength; got != 1024 {
		t.Fatalf("Unexpected MaxControlFramePayloadSize: %d", got)
	}
	if got := framer.maxHeaderFieldSize; got != 128 {
		t.Fatalf("Unexpected MaxHeaderFieldSize: %d", got)
	}
	if got := framer.maxHeaderCount; got != 16 {
		t.Fatalf("Unexpected MaxHeaderCount: %d", got)
	}
}
