package base64encdec

import (
	b64 "encoding/base64"
)

// EncodeString encodes base64 string
func EncodeString(str string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	return sEnc
}
