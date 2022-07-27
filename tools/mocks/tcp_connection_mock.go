package mocks

import (
	"errors"
	"mtt/structures"
)

type MockTCPConnection struct {
	Success bool
}

func (c *MockTCPConnection) Connect() error {
	return nil
}

func (c *MockTCPConnection) Send(data []byte) error {
	if len(data) == 0 {
		return errors.New("empty data")
	}
	return nil
}

func (c *MockTCPConnection) Get() ([]byte, error) {
	var response structures.Response
	if c.Success {
		header := structures.Header{
			SvcId:      1,
			BodyLength: 1024,
			RequestId:  1,
		}
		response = structures.Response{
			Header: header,
		}

		clientId := new(structures.String)
		clientId.FromString("test_client_id")
		username := new(structures.String)
		username.FromString("test_username")

		response.Code = structures.CODE_OK
		response.Body = &structures.SvcOkResponseBody{
			ClientId:   *clientId,
			ClientType: 13,
			Username:   *username,
			ExpiresIn:  1296000,
			UserId:     42,
		}
	} else {
		message := new(structures.String)
		message.FromString("bad scope")

		response.Code = structures.CODE_BAD_SCOPE
		response.Body = &structures.SvcErrorResponseBody{
			Code:        structures.CODE_BAD_SCOPE,
			ErrorString: *message,
		}
	}

	encodedResponse, err := response.Encode()
	if err != nil {
		return nil, err
	}

	return encodedResponse, nil
}

func (c *MockTCPConnection) Close() (err error) {
	return
}
