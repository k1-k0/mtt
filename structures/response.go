package structures

import (
	"bytes"
	"encoding/binary"
)

const (
	CODE_OK              = 0x00000000
	CODE_TOKEN_NOT_FOUND = 0x00000001
	CODE_DB_ERROR        = 0x00000002
	CODE_UNKNOWN_MSG     = 0x00000003
	CODE_BAD_PACKET      = 0x00000004
	CODE_BAD_CLIENT      = 0x00000005
	CODE_BAD_SCOPE       = 0x00000006
)

var codeInfo = map[int]string{
	CODE_OK:              "CUBE_OAUTH2_ERR_OK",
	CODE_TOKEN_NOT_FOUND: "CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND",
	CODE_DB_ERROR:        "CUBE_OAUTH2_ERR_DB_ERROR",
	CODE_UNKNOWN_MSG:     "CUBE_OAUTH2_ERR_UNKNOWN_MSG",
	CODE_BAD_PACKET:      "CUBE_OAUTH2_ERR_BAD_PACKET",
	CODE_BAD_CLIENT:      "CUBE_OAUTH2_ERR_BAD_CLIENT",
	CODE_BAD_SCOPE:       "CUBE_OAUTH2_ERR_BAD_SCOPE",
}

type Response struct {
	Header Header
	Code   ReturnCode
	Body   ResponseBody
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

	returnCodeData, err := r.Code.Encode()
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

	err = r.Code.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"return_code"}
	}

	buffer.Next(r.Code.GetSizeOfBytes())

	if r.Code == CODE_OK {
		r.Body = &SvcOkResponseBody{}
	} else {
		r.Body = &SvcErrorResponseBody{Code: r.Code}
	}

	err = r.Body.Decode(buffer.Bytes())
	if err != nil {
		return &InvalidDecodingError{"body"}
	}

	return nil
}

func (r *Response) String() string {
	if r.Code == CODE_OK {
		return r.Body.(*SvcOkResponseBody).String()
	} else {
		return r.Body.(*SvcErrorResponseBody).String()
	}
}

func (r *Response) Equal(body *Response) bool {
	if !r.Header.Equal(&body.Header) {
		return false
	}

	if r.Code != body.Code {
		return false
	}

	switch bd := r.Body.(type) {
	case *SvcOkResponseBody:
		if !r.Body.(*SvcOkResponseBody).Equal(bd) {
			return false
		}
	case *SvcErrorResponseBody:
		if !r.Body.(*SvcErrorResponseBody).Equal(bd) {
			return false
		}
	}

	return true
}
