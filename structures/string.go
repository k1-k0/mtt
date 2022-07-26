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

func (s String) GetSizeOfBytes() int {
	return int(unsafe.Sizeof(s.Length)) + len(s.Str)
}

func (s *String) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	err := binary.Write(data, binary.LittleEndian, &s.Length)
	if err != nil {
		return nil, &InvalidEncodingError{"lenght"}
	}

	for i := range s.Str {
		err = data.WriteByte(byte(s.Str[i]))
		if err != nil {
			errMsg := fmt.Sprintf("invalid encoding of byte #%d", i)
			return nil, errors.New(errMsg)
		}
	}

	return data.Bytes(), nil
}

func (s *String) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	if err := binary.Read(buffer, binary.LittleEndian, &s.Length); err != nil {
		return &InvalidDecodingError{"lenght"}
	}

	s.Str = make([]int8, s.Length)
	for i := int32(0); i < s.Length; i++ {
		err := binary.Read(buffer, binary.LittleEndian, &s.Str[i])
		if err != nil {
			errMsg := fmt.Sprintf("invalid decoding of byte #%d", i)
			return errors.New(errMsg)
		}
	}

	return nil
}

func (s *String) FromString(str string) {
	s.Length = int32(len(str))
	s.Str = make([]int8, 0, s.Length)

	for _, ch := range str {
		s.Str = append(s.Str, int8(ch))
	}
}

func (s String) String() string {
	builder := new(strings.Builder)

	for i := range s.Str {
		err := builder.WriteByte(byte(s.Str[i]))
		if err != nil {
			fmt.Printf("invalid decoding of byte #%d", i)
			return ""
		}
	}

	return builder.String()
}

func (s *String) Equal(str *String) bool {
	if (len(s.Str) != len(str.Str)) || (s.Length != str.Length) {
		return false
	}

	for i := range s.Str {
		if s.Str[i] != str.Str[i] {
			return false
		}
	}

	return true
}
