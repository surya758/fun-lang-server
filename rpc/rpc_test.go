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
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	if contentLength != 15 {
		t.Fatalf("Expected 16, Got %d", contentLength)
	}
	if method != "hi" {
		t.Fatalf("Expected hi, Got %s", method)
	}
}
