package structures

import (
	"strings"
	"testing"
)

func makeResponse(code int) Response {
	header := Header{1, 1024, 1}
	response := Response{
		Header: header,
	}

	if code == CODE_OK {
		clientId := new(String)
		clientId.FromString("test_client_id")
		username := new(String)
		username.FromString("test_username")

		response.Code = CODE_OK
		response.Body = &SvcOkResponseBody{
			ClientId:   *clientId,
			ClientType: 13,
			Username:   *username,
			ExpiresIn:  1296000,
			UserId:     42,
		}
	} else {
		message := new(String)
		message.FromString("bad packet")

		response.Code = CODE_BAD_PACKET
		response.Body = &SvcErrorResponseBody{
			Code:        CODE_BAD_PACKET,
			ErrorString: *message,
		}
	}

	return response
}

func TestOkResponseEncodeDecode(t *testing.T) {
	response := makeResponse(CODE_OK)

	encodedResponse, err := response.Encode()
	if err != nil {
		t.Fatalf("invalid response encoding. %s", err.Error())
	}

	var decodedResponse Response
	err = decodedResponse.Decode(encodedResponse)
	if err != nil {
		t.Fatalf("invalid response decoding. %s", err.Error())
	}

	if response.Equal(&decodedResponse) == false {
		t.Fatalf("response objects are not equal. expected %v but actual %v",
			response, decodedResponse)
	}
}

func TestOkResponseEqual(t *testing.T) {
	response := makeResponse(CODE_OK)

	if response.Equal(&response) == false {
		t.Fatalf("svc_ok_response_body objects are not equal. expected %v but actual %v",
			response, response)
	}
}

func TestOkResponseConvertationToString(t *testing.T) {
	response := makeResponse(CODE_OK)

	expectedString := "client_id: test_client_id\nclient_type: 13\nexpires_in: 1296000\nuser_id: 42\nusername: test_username\n"

	actualString := response.Body.(*SvcOkResponseBody).String()

	if strings.Compare(expectedString, actualString) != 0 {
		t.Fatalf("strings are not equal. expected: %s but actual: %s", expectedString, actualString)
	}
}

func TestErrorResponseEncodeDecode(t *testing.T) {
	response := makeResponse(CODE_BAD_PACKET)

	encodedResponse, err := response.Encode()
	if err != nil {
		t.Fatalf("invalid response encoding. %s", err.Error())
	}

	var decodedResponse Response
	err = decodedResponse.Decode(encodedResponse)
	if err != nil {
		t.Fatalf("invalid response decoding. %s", err.Error())
	}

	if response.Equal(&decodedResponse) == false {
		t.Fatalf("response objects are not equal. expected %v but actual %v",
			response, decodedResponse)
	}
}

func TestErrorResponseEqual(t *testing.T) {
	response := makeResponse(CODE_BAD_PACKET)

	if response.Equal(&response) == false {
		t.Fatalf("svc_ok_response_body objects are not equal. expected %v but actual %v",
			response, response)
	}
}

func TestErrorResponseConvertationToString(t *testing.T) {
	response := makeResponse(CODE_BAD_PACKET)

	expectedString := "error: CUBE_OAUTH2_ERR_BAD_PACKET\nmessage: bad packet\n"

	actualString := response.Body.(*SvcErrorResponseBody).String()

	if strings.Compare(expectedString, actualString) != 0 {
		t.Fatalf("strings are not equal. expected: %s but actual: %s", expectedString, actualString)
	}
}
