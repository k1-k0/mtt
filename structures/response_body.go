package structures

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type ResponseBody interface {
	Encode() ([]byte, error)
	Decode(data []byte) error
}

type ReturnCode int32

type SvcOkResponseBody struct {
	ClientId   String
	ClientType int32
	Username   String
	ExpiresIn  int32
	UserId     int64
}

type SvcErrorResponseBody struct {
	Code        ReturnCode
	ErrorString String
}

func (r *ReturnCode) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, r)
	if err != nil {
		return nil, &InvalidEncodingError{"return_code"}
	}

	return data.Bytes(), nil
}

func (r *ReturnCode) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := binary.Read(buffer, binary.LittleEndian, r)
	if err != nil {
		return &InvalidDecodingError{"return_code"}
	}

	return nil
}

func (b *SvcOkResponseBody) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	clientIdData, err := b.ClientId.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"client_id"}
	}

	err = binary.Write(data, binary.LittleEndian, clientIdData)
	if err != nil {
		return nil, &InvalidEncodingError{"client_type"}
	}

	err = binary.Write(data, binary.LittleEndian, b.ClientType)
	if err != nil {
		return nil, &InvalidEncodingError{"client_type"}
	}

	usernameData, err := b.Username.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"username"}
	}

	err = binary.Write(data, binary.LittleEndian, usernameData)
	if err != nil {
		return nil, &InvalidEncodingError{"username"}
	}

	err = binary.Write(data, binary.LittleEndian, b.ExpiresIn)
	if err != nil {
		return nil, &InvalidEncodingError{"expires_in"}
	}

	err = binary.Write(data, binary.LittleEndian, b.UserId)
	if err != nil {
		return nil, &InvalidEncodingError{"user_id"}
	}

	return data.Bytes(), nil
}

func (b *SvcOkResponseBody) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := b.ClientId.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"client_id"}
	}

	buffer.Next(b.ClientId.GetSizeOfBytes())

	err = binary.Read(buffer, binary.LittleEndian, &b.ClientType)
	if err != nil {
		return &InvalidDecodingError{"client_type"}
	}

	err = b.Username.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"username"}
	}

	buffer.Next(b.Username.GetSizeOfBytes())

	err = binary.Read(buffer, binary.LittleEndian, &b.ExpiresIn)
	if err != nil {
		return &InvalidDecodingError{"expires_in"}
	}

	err = binary.Read(buffer, binary.LittleEndian, &b.UserId)
	if err != nil {
		return &InvalidDecodingError{"user_id"}
	}

	return nil
}

func (b SvcOkResponseBody) String() string {
	format := "client_id: %s\nclient_type: %d\nexpires_in: %d\nuser_id: %d\nusername: %s\n"
	return fmt.Sprintf(format, b.ClientId, b.ClientType, b.ExpiresIn, b.UserId, b.Username)
}

func (b *SvcOkResponseBody) Equal(body *SvcOkResponseBody) bool {
	if !b.ClientId.Equal(&body.ClientId) {
		return false
	}

	if b.ClientType != body.ClientType {
		return false
	}

	if !b.Username.Equal(&body.Username) {
		return false
	}

	if b.ExpiresIn != body.ExpiresIn {
		return false
	}

	if b.UserId != body.UserId {
		return false
	}

	return true
}

func (b *SvcErrorResponseBody) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, b.Code)
	if err != nil {
		return nil, &InvalidEncodingError{"code"}
	}

	clientIdData, err := b.ErrorString.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"error_string"}
	}

	err = binary.Write(data, binary.LittleEndian, clientIdData)
	if err != nil {
		return nil, &InvalidEncodingError{"error_string"}
	}

	return data.Bytes(), nil
}

func (b *SvcErrorResponseBody) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := binary.Read(buffer, binary.LittleEndian, &b.Code)
	if err != nil {
		return &InvalidDecodingError{"client_type"}
	}

	err = b.ErrorString.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"error_string"}
	}

	return nil
}

func (b SvcErrorResponseBody) String() string {
	format := "error: %s\nmessage: %s\n"
	return fmt.Sprintf(format, codeInfo[int(b.Code)], b.ErrorString)
}

func (b *SvcErrorResponseBody) Equal(body *SvcErrorResponseBody) bool {
	if b.Code != body.Code {
		return false
	}

	if !b.ErrorString.Equal(&body.ErrorString) {
		return false
	}

	return true
}
