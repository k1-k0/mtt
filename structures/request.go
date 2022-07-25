package structures

import (
	"bytes"
	"encoding/binary"
)

type Request struct {
	Header Header
	Body   SvcRequestBody
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
