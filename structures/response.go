package structures

import (
	"bytes"
	"encoding/binary"
)

type Response struct {
	Header     Header
	ReturnCode ResponseReturnCode
	Body       ResponseBody
}

func (r *Response) Encode() ([]byte, error) {
	data := new(bytes.Buffer)

	headerData, err := r.Header.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"header"}
	}

	err = binary.Write(data, binary.LittleEndian, headerData)
	if err != nil {
		return nil, &InvalidEncodingError{"header"}
	}

	returnCodeData, err := r.ReturnCode.Encode()
	if err != nil {
		return nil, &InvalidEncodingError{"return_code"}
	}

	err = binary.Write(data, binary.LittleEndian, returnCodeData)
	if err != nil {
		return nil, &InvalidEncodingError{"return_code"}
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

func (r *Response) Decode(data []byte) error {
	buffer := bytes.NewBuffer(data)

	err := r.Header.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"header"}
	}

	buffer.Next(r.Header.GetSizeOfBytes())

	err = r.ReturnCode.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"return_code"}
	}

	if r.ReturnCode == OK_CODE {
		r.Body = &SvcOkResponseBody{}
	} else {
		r.Body = &SvcErrorResponseBody{}
	}

	err = r.Body.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"body"}
	}

	return nil
}
