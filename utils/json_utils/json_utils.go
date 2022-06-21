package json_utils

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"my_qor_test/config/consts"
	"regexp"
)

func Decode(data []byte, v interface{}) error {
	str := string(data)
	return DecodeString(str, v)
}

func DecodeString(str string, v interface{}) error {
	re := regexp.MustCompile(consts.Expr4)
	str = re.ReplaceAllStringFunc(str, func(s string) string {
		if s == `\r` {
			return `\\r`
		}
		if s == `\n` {
			return `\\n`
		}
		return s
	})
	data := []byte(str)
	if err := json.Unmarshal(data, v); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func Encode(v interface{}) ([]byte, error) {
	return BaseEncode(v, false, true)
}
func EncodeString(v interface{}) (string, error) {
	data, err := Encode(v)
	return string(data), err
}

func BaseEncode(v interface{}, escapeHTML bool, escapeLF bool) ([]byte, error) {
	if escapeHTML {
		marshaledBytes, err := json.Marshal(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if !escapeLF {
			marshaledBytes = append(marshaledBytes, 10)
		}
		return marshaledBytes, nil
	} else {
		encodeBuffer := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(encodeBuffer)
		encoder.SetEscapeHTML(escapeHTML)
		if err := encoder.Encode(v); err != nil {
			return nil, errors.WithStack(err)
		}
		encodedBytes := encodeBuffer.Bytes()
		if escapeLF {
			if len(encodedBytes) > 0 {
				dataLength := len(encodedBytes)
				if encodedBytes[dataLength-1] == 10 {
					encodedBytes = encodedBytes[0 : dataLength-1]
				}
			}
		}
		return encodedBytes, nil
	}
}
