package structures

import (
	"testing"
)

func TestRequestEncodeDecode(t *testing.T) {
	var (
		token = "test"
		scope = "testing"
	)

	request := Request{}
	requestBytes, err := request.BuildRequest(token, scope)
	if err != nil {
		t.Fatalf("invalid request building. %s", err.Error())
	}

	var decodedRequest Request
	err = decodedRequest.Decode(requestBytes)
	if err != nil {
		t.Fatalf("invalid request decoding. %s", err.Error())
	}

	encodedRequest, err := decodedRequest.Encode()
	if err != nil {
		t.Fatalf("invalid request encoding. %s", err.Error())
	}

	for i := range encodedRequest {
		if encodedRequest[i] != requestBytes[i] {
			t.Fatalf("request bytes are not equal. expected %v but actual %v",
				requestBytes[i], encodedRequest[i])
		}
	}
}

func TestRequestEncodeDecodeEmptyData(t *testing.T) {
	var (
		token = ""
		scope = ""
	)

	request := Request{}
	requestBytes, err := request.BuildRequest(token, scope)
	if err != nil {
		t.Fatalf("invalid request building. %s", err.Error())
	}

	var decodedRequest Request
	err = decodedRequest.Decode(requestBytes)
	if err != nil {
		t.Fatalf("invalid request decoding. %s", err.Error())
	}

	encodedRequest, err := decodedRequest.Encode()
	if err != nil {
		t.Fatalf("invalid request encoding. %s", err.Error())
	}

	for i := range encodedRequest {
		if encodedRequest[i] != requestBytes[i] {
			t.Fatalf("request bytes are not equal. expected %v but actual %v",
				requestBytes[i], encodedRequest[i])
		}
	}
}
