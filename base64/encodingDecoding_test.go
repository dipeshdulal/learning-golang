package base64encdec_test

import (
	"testing"

	base64encdec "wesionary.team/dipeshdulal/golang-test/base64"
)

func TestBase64Encoding(t *testing.T) {
	d := base64encdec.EncodeString("Hello नेपाल")
	t.Logf("Data returned is after ENCODING %v", d)
	origString, _ := base64encdec.DecodeString(d)
	t.Logf("Data returned is after DECODING %v", origString)
}
