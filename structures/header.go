package structures

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type Header struct {
	SvcId      int32
	BodyLength int32
	RequestId  int32
}

func (h *Header) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, h)
	if err != nil {
		// TODO: Make separate error
		return nil, errors.New("Invalid encoding of header")
	}

	return data.Bytes(), nil
}

func (h *Header) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := binary.Read(buffer, binary.LittleEndian, h)
	if err != nil {
		// TODO: Make separate error
		return errors.New("Invalid decoding of header")
	}

	return nil
}
