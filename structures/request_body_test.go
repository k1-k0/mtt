package structures

import (
	"testing"
)

func TestSvcRequestBodyEncodeDecode(t *testing.T) {
	token := new(String)
	token.FromString("test")

	scope := new(String)
	scope.FromString("testing")

	svcRequestBody := SvcRequestBody{1, *token, *scope}

	encodedSvcRequestBody, err := svcRequestBody.Encode()
	if err != nil {
		t.Fatalf("invalid svc_request_body encoding. %s", err.Error())
	}

	var decodedSvcRequestBody SvcRequestBody
	err = decodedSvcRequestBody.Decode(encodedSvcRequestBody)
	if err != nil {
		t.Fatalf("invalid svc_request_body decoding. %s", err.Error())
	}

	if svcRequestBody.Equal(&decodedSvcRequestBody) == false {
		t.Fatalf("svc_request_body objects are not equal. expected %v but actual %v",
			svcRequestBody, decodedSvcRequestBody)
	}
}

func TestSvcRequestBodyEncodeDecodeEmpty(t *testing.T) {
	svcRequestBody := SvcRequestBody{}

	encodedSvcRequestBody, err := svcRequestBody.Encode()
	if err != nil {
		t.Fatalf("invalid svc_request_body encoding. %s", err.Error())
	}

	var decodedSvcRequestBody SvcRequestBody
	err = decodedSvcRequestBody.Decode(encodedSvcRequestBody)
	if err != nil {
		t.Fatalf("invalid svc_request_body decoding. %s", err.Error())
	}

	if svcRequestBody.Equal(&decodedSvcRequestBody) == false {
		t.Fatalf("svc_request_body objects are not equal. expected %v but actual %v",
			svcRequestBody, decodedSvcRequestBody)
	}
}

func TestSvcRequestBodyEqual(t *testing.T) {
	token := new(String)
	token.FromString("main")
	scope := new(String)
	scope.FromString("production")

	body := SvcRequestBody{4, *token, *scope}

	if body.Equal(&body) == false {
		t.Fatalf("svc_request_body objects are not equal. expected %v but actual %v", body, body)
	}
}
