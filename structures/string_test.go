package structures

import (
	"strings"
	"testing"
)

func StringToOwnString(str string) String {
	s := new(String)
	s.FromString(str)
	return *s
}

func TestStringGetSize(t *testing.T) {
	s := new(String)
	s.FromString("test")
	expectedLength := 8

	actualLength := s.GetSizeOfBytes()

	if expectedLength != actualLength {
		t.Fatalf("length are not equal. expected %d but actual %d", expectedLength, actualLength)
	}
}

func TestStringEncodeDecode(t *testing.T) {
	data := "encode_decode"
	expectedString := StringToOwnString(data)

	encodedString, err := expectedString.Encode()
	if err != nil {
		t.Fatalf("invalid string encoding. %s", err.Error())
	}

	var decodedString String
	err = decodedString.Decode(encodedString)
	if err != nil {
		t.Fatalf("invalid string decoding. %s", err.Error())
	}

	if expectedString.Equal(&decodedString) == false {
		t.Fatalf("strings are not equal. expected %v but actual %v", expectedString, decodedString)
	}
}

func TestStringEncodeDecodeEmpty(t *testing.T) {
	data := ""
	expectedString := StringToOwnString(data)

	encodedString, err := expectedString.Encode()
	if err != nil {
		t.Fatalf("invalid string encoding. %s", err.Error())
	}

	var decodedString String
	err = decodedString.Decode(encodedString)
	if err != nil {
		t.Fatalf("invalid string decoding. %s", err.Error())
	}

	if expectedString.Equal(&decodedString) == false {
		t.Fatalf("strings are not equal. expected %v but actual %v", expectedString, decodedString)
	}
}

func TestStringFromString(t *testing.T) {
	data := "test"

	actualString := new(String)
	expectedBytes := make([]int8, len(data))
	for i := range data {
		expectedBytes = append(expectedBytes, int8(data[i]))
	}
	expectedString := String{
		Length: int32(4),
		Str:    expectedBytes,
	}

	actualString.FromString(data)

	if expectedString.Equal(actualString) {
		t.Fatalf("strings are not equal. expected %v but actual %v", expectedString, actualString)
	}
}

func TestStringConvertationToString(t *testing.T) {
	data := "test"
	expectedData := data

	actualString := new(String)
	actualString.FromString(data)

	actualData := actualString.String()

	if strings.Compare(expectedData, actualData) != 0 {
		t.Fatalf("strings are not equal. expected %v but actual %v", expectedData, actualData)
	}
}

func TestStringEqual(t *testing.T) {
	data := "abracadabra"
	s := new(String)
	s.FromString(data)

	if s.Equal(s) == false {
		t.Fatalf("strings are not equal. expected %v but actual %v", s, s)
	}
}
