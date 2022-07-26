package structures

import (
	"strings"
	"testing"
)

func makeOkResponse() *SvcOkResponseBody {

	clientId := new(String)
	clientId.FromString("test_client_id")
	username := new(String)
	username.FromString("test_username")

	return &SvcOkResponseBody{
		ClientId:   *clientId,
		ClientType: 13,
		Username:   *username,
		ExpiresIn:  1296000,
		UserId:     42,
	}
}

func TestResponseOkEncodeDecode(t *testing.T) {
	svcOkResponseBody := makeOkResponse()

	encodedSvcOkResponseBody, err := svcOkResponseBody.Encode()
	if err != nil {
		t.Fatalf("invalid svc_ok_response_body encoding. %s", err.Error())
	}

	var decodedSvcOkResponseBody SvcOkResponseBody
	err = decodedSvcOkResponseBody.Decode(encodedSvcOkResponseBody)
	if err != nil {
		t.Fatalf("invalid svc_ok_response_body decoding. %s", err.Error())
	}

	if svcOkResponseBody.Equal(&decodedSvcOkResponseBody) == false {
		t.Fatalf("svc_ok_response_body objects are not equal. expected %v but actual %v",
			svcOkResponseBody, decodedSvcOkResponseBody)
	}
}

func TestResponseOkEncodeDecodeEmpty(t *testing.T) {
	svcOkResponseBody := SvcOkResponseBody{}

	encodedSvcOkResponseBody, err := svcOkResponseBody.Encode()
	if err != nil {
		t.Fatalf("invalid svc_ok_response_body encoding. %s", err.Error())
	}

	var decodedSvcOkResponseBody SvcOkResponseBody
	err = decodedSvcOkResponseBody.Decode(encodedSvcOkResponseBody)
	if err != nil {
		t.Fatalf("invalid svc_ok_response_body decoding. %s", err.Error())
	}

	if svcOkResponseBody.Equal(&decodedSvcOkResponseBody) == false {
		t.Fatalf("svc_ok_response_body objects are not equal. expected %v but actual %v",
			svcOkResponseBody, decodedSvcOkResponseBody)
	}
}

func TestResponseOkEqual(t *testing.T) {
	svcOkResponseBody := makeOkResponse()

	if svcOkResponseBody.Equal(svcOkResponseBody) == false {
		t.Fatalf("svc_ok_response_body objects are not equal. expected %v but actual %v",
			svcOkResponseBody, svcOkResponseBody)
	}
}

func TestSvcResponseOkConvertationToString(t *testing.T) {
	svcOkResponseBody := makeOkResponse()
	expectedString := "client_id: test_client_id\nclient_type: 13\nexpires_in: 1296000\nuser_id: 42\nusername: test_username\n"

	actualString := svcOkResponseBody.String()

	if strings.Compare(expectedString, actualString) != 0 {
		t.Fatalf("strings are not equal. expected: %s but actual: %s", expectedString, actualString)
	}
}

func TestResponseErrorEncodeDecode(t *testing.T) {
	message := new(String)
	message.FromString("token not found")

	svcErrorResponseBody := SvcErrorResponseBody{
		Code:        CODE_TOKEN_NOT_FOUND,
		ErrorString: *message,
	}

	encodedSvcErrorResponseBody, err := svcErrorResponseBody.Encode()
	if err != nil {
		t.Fatalf("invalid svc_error_response_body encoding. %s", err.Error())
	}

	var decodedSvcErrorResponseBody SvcErrorResponseBody
	err = decodedSvcErrorResponseBody.Decode(encodedSvcErrorResponseBody)
	if err != nil {
		t.Fatalf("invalid svc_error_response_body decoding. %s", err.Error())
	}

	if svcErrorResponseBody.Equal(&decodedSvcErrorResponseBody) == false {
		t.Fatalf("svc_error_response_body objects are not equal. expected %v but actual %v",
			svcErrorResponseBody, decodedSvcErrorResponseBody)
	}
}

func TestResponseErrorEncodeDecodeEmpty(t *testing.T) {
	svcErrorResponseBody := SvcErrorResponseBody{}

	encodedSvcErrorResponseBody, err := svcErrorResponseBody.Encode()
	if err != nil {
		t.Fatalf("invalid svc_error_response_body encoding. %s", err.Error())
	}

	var decodedSvcErrorResponseBody SvcErrorResponseBody
	err = decodedSvcErrorResponseBody.Decode(encodedSvcErrorResponseBody)
	if err != nil {
		t.Fatalf("invalid svc_error_response_body decoding. %s", err.Error())
	}

	if svcErrorResponseBody.Equal(&decodedSvcErrorResponseBody) == false {
		t.Fatalf("svc_error_response_body objects are not equal. expected %v but actual %v",
			svcErrorResponseBody, decodedSvcErrorResponseBody)
	}
}

func TestResponseErrorEqual(t *testing.T) {
	message := new(String)
	message.FromString("token not found")

	svcErrorResponseBody := SvcErrorResponseBody{
		Code:        CODE_TOKEN_NOT_FOUND,
		ErrorString: *message,
	}

	if svcErrorResponseBody.Equal(&svcErrorResponseBody) == false {
		t.Fatalf("svc_ok_response_body objects are not equal. expected %v but actual %v",
			svcErrorResponseBody, svcErrorResponseBody)
	}
}

func TestSvcResponseErrorConvertationToString(t *testing.T) {
	message := new(String)
	message.FromString("token not found")
	svcErrorResponseBody := SvcErrorResponseBody{
		Code:        CODE_TOKEN_NOT_FOUND,
		ErrorString: *message,
	}
	expectedString := "error: CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND\nmessage: token not found\n"

	actualString := svcErrorResponseBody.String()

	if strings.Compare(expectedString, actualString) != 0 {
		t.Fatalf("strings are not equal. expected %s but actual %s", expectedString, actualString)
	}
}
