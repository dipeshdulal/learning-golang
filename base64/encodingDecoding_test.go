package base64encdec_test

import (
	"testing"

	base64encdec "wesionary.team/dipeshdulal/golang-test/base64"
)

func TestBase64Encoding(t *testing.T) {
	d := base64encdec.EncodeString("Hello World")
	t.Logf("Data returned is: %v", d)
}
