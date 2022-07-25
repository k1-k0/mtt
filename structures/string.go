package structures

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

type String struct {
	Length int32
	Str    []int8
}

func (s String) GetLength() int {
	return int(unsafe.Sizeof(s.Length)) + len(s.Str)
}

func (s *String) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, &s.Length)
	if err != nil {
		// TODO: Make separate error
		return nil, errors.New("Invalid encoding of lenght of string")
	}

	for i := range s.Str {
		err = data.WriteByte(byte(s.Str[i]))
		if err != nil {
			// TODO: Make separate error
			errMsg := fmt.Sprintf("Invalid encoding of byte #%d of string", i)
			return nil, errors.New(errMsg)
		}
	}

	return data.Bytes(), nil
}

func (s *String) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	if err := binary.Read(buffer, binary.LittleEndian, &s.Length); err != nil {
		// TODO: Make separate error
		return errors.New("Invalid decoding of lenght of string")
	}

	s.Str = make([]int8, s.Length, s.Length)
	for i := int32(0); i < s.Length; i++ {
		err := binary.Read(buffer, binary.LittleEndian, &s.Str[i])
		if err != nil {
			// TODO: Make separate error
			errMsg := fmt.Sprintf("Invalid decoding of byte #%d of string", i)
			return errors.New(errMsg)
		}
	}

	return nil
}

func (s String) String() string {
	builder := new(strings.Builder)

	for i := range s.Str {
		err := builder.WriteByte(byte(s.Str[i]))
		if err != nil {
			fmt.Printf("Invalid decoding of byte #%d string", i)
			return ""
		}
	}

	return builder.String()
}
