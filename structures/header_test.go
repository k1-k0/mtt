package structures

import (
	"testing"
)

func TestHeaderGetSize(t *testing.T) {
	h := Header{1, 2, 3}
	expectedLength := 4 * 3

	actualLength := h.GetSizeOfBytes()

	if expectedLength != actualLength {
		t.Fatalf("length are not equal. expected %d but actual %d", expectedLength, actualLength)
	}
}

func TestHeaderEncodeDecode(t *testing.T) {
	expectedHeader := Header{1, 1, 1}

	encodedHeader, err := expectedHeader.Encode()
	if err != nil {
		t.Fatalf("invalid header encoding. %s", err.Error())
	}

	var decodedHeader Header
	err = decodedHeader.Decode(encodedHeader)
	if err != nil {
		t.Fatalf("invalid header decoding. %s", err.Error())
	}

	if expectedHeader.Equal(&decodedHeader) == false {
		t.Fatalf("headers are not equal. expected %v but actual %v", expectedHeader, decodedHeader)
	}
}

func TestHeaderEqual(t *testing.T) {
	h := Header{1, 1, 1}

	if h.Equal(&h) == false {
		t.Fatalf("headers are not equal. expected %v but actual %v", h, h)
	}
}
