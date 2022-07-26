package structures

import (
	"bytes"
	"encoding/binary"
)

const (
	OK_CODE = 0
)

type ResponseBody interface {
	Encode() ([]byte, error)
	Decode(data []byte) error
}

// ResponseError when return_code != 0, otherwise Ok
type ResponseReturnCode int32

type SvcOkResponseBody struct {
	ClientId   String
	ClientType int32
	Username   String
	ExpiresIn  int32
	UserId     int64
}

type SvcErrorResponseBody struct {
	ErrorString String
}

func (r *ResponseReturnCode) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, r)
	if err != nil {
		return nil, &InvalidEncodingError{"return_code"}
	}

	return data.Bytes(), nil
}

func (r *ResponseReturnCode) Decode(data []byte) error {
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
		return nil, &InvalidEncodingError{"scope"}
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
		return nil, &InvalidEncodingError{"scope"}
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

	err = binary.Read(buffer, binary.LittleEndian, b.ClientType)
	if err != nil {
		return &InvalidDecodingError{"client_type"}
	}

	err = b.Username.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"username"}
	}

	buffer.Next(b.Username.GetSizeOfBytes())

	err = binary.Read(buffer, binary.LittleEndian, b.ExpiresIn)
	if err != nil {
		return &InvalidDecodingError{"expires_in"}
	}

	err = binary.Read(buffer, binary.LittleEndian, b.UserId)
	if err != nil {
		return &InvalidDecodingError{"user_id"}
	}

	return nil
}

func (b *SvcErrorResponseBody) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := b.ErrorString.Decode(data.Bytes())
	if err != nil {
		return nil, &InvalidEncodingError{"error_string"}
	}

	return data.Bytes(), nil
}

func (b *SvcErrorResponseBody) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := b.ErrorString.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"error_string"}
	}

	return nil
}
