package structures

import "fmt"

type InvalidEncodingError struct {
	What string
}

type InvalidDecodingError struct {
	What string
}

func (e *InvalidEncodingError) Error() string {
	return fmt.Sprintf("invalid encoding of '%s'", e.What)
}

func (e *InvalidDecodingError) Error() string {
	return fmt.Sprintf("invalid decoding of '%s'", e.What)
}
