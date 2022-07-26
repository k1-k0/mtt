package structures

import (
	"bytes"
	"encoding/binary"
)

const (
	SVC_MSG    = 0x00000001
	SVC_ID     = 0x00000002
	REQUEST_ID = 0x00000001
)

type Request struct {
	Header Header
	Body   SvcRequestBody
}

func (r *Request) BuildRequest(token, scope string) ([]byte, error) {

	data := new(bytes.Buffer)

	tokenString := String{}
	tokenString.FromString(token)

	scopeString := String{}
	scopeString.FromString(scope)

	r.Body = SvcRequestBody{
		SvcMsg: SVC_MSG,
		Token:  tokenString,
		Scope:  scopeString,
	}

	bodyData, err := r.Body.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"request body"}
	}

	err = binary.Write(data, binary.LittleEndian, bodyData)
	if err != nil {
		return nil, &InvalidEncodingError{"request body"}
	}

	r.Header = Header{
		SvcId:      SVC_ID,
		BodyLength: int32(len(bodyData)),
		RequestId:  REQUEST_ID,
	}

	headerData, err := r.Header.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"request header"}
	}

	err = binary.Write(data, binary.LittleEndian, headerData)
	if err != nil {
		return nil, &InvalidEncodingError{"request header"}
	}

	return data.Bytes(), nil
}

func (r *Request) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	headerData, err := r.Header.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"header"}
	}

	err = binary.Write(data, binary.LittleEndian, headerData)
	if err != nil {
		return nil, &InvalidEncodingError{"header"}
	}

	bodyData, err := r.Body.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"body"}
	}

	err = binary.Write(data, binary.LittleEndian, bodyData)
	if err != nil {
		return nil, &InvalidEncodingError{"body"}
	}

	return data.Bytes(), nil
}

func (r *Request) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := r.Header.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"header"}
	}

	buffer.Next(r.Header.GetSizeOfBytes())

	err = r.Body.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"body"}
	}

	return nil
}
