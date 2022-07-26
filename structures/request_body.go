package structures

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type SvcRequestBody struct {
	SvcMsg int32
	Token  String
	Scope  String
}

func (b *SvcRequestBody) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, b.SvcMsg)
	if err != nil {
		return nil, &InvalidEncodingError{"svc_msg"}
	}

	tokenBytes, err := b.Token.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"scope"}
	}

	err = binary.Write(data, binary.LittleEndian, tokenBytes)
	if err != nil {
		return nil, &InvalidEncodingError{"scope"}
	}

	scopeBytes, err := b.Scope.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"scope"}
	}

	err = binary.Write(data, binary.LittleEndian, scopeBytes)
	if err != nil {
		return nil, &InvalidEncodingError{"scope"}
	}

	return data.Bytes(), nil
}

func (b *SvcRequestBody) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := binary.Read(buffer, binary.LittleEndian, b.SvcMsg)
	if err != nil {
		return &InvalidDecodingError{"svc_msg"}
	}

	err = binary.Read(buffer, binary.LittleEndian, b.Token)
	if err != nil {
		return errors.New("invalid read bytes of token")
	}

	err = b.Token.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"token"}
	}

	buffer.Next(b.Token.GetSizeOfBytes())

	err = b.Scope.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"scope"}
	}

	return nil
}
