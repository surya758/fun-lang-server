package rpc_test

import (
	"testing"

	"github.com/surya758/fun-lang-server/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncoding(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected != actual {
		t.Fatalf("Expected %s, got %s", expected, actual)
	}
}

func TestDecoding(t *testing.T) {
	incomingMessage := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if contentLength != 16 {
		t.Fatalf("Expected 16, Got %d", contentLength)
	}
}
