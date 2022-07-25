package structures

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

type Header struct {
	SvcId      int32
	BodyLength int32
	RequestId  int32
}

func (h Header) GetSizeOfBytes() int {
	return int(unsafe.Sizeof(h.SvcId)) * 3
}

func (h *Header) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, h)
	if err != nil {
		return nil, &InvalidEncodingError{"header"}
	}

	return data.Bytes(), nil
}

func (h *Header) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := binary.Read(buffer, binary.LittleEndian, h)
	if err != nil {
		return &InvalidDecodingError{"header"}
	}

	return nil
}
