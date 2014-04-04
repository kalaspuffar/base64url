package base64url

import (
	"github.com/kalaspuffar/base64url"
	"testing"
	"encoding/hex"
)

var (
	HEXEncoded = "69b73efeba"
	Base64URLEncoded = "abc-_ro"
)

/*
	Test base64url.Encode function using a hex encoded string corresponding to an Base64URLEncoded string.
	First we decode the hex encoded bytes and then run the encode function on this and check the result with the 
	corresponding Base64url encoded data.
 */
func TestEncode(t *testing.T) {
	testBytes := make([]byte, len(HEXEncoded) / 2)
	hex.Decode(testBytes, []byte(HEXEncoded))
	resultStr := base64url.Encode(testBytes)
	if(Base64URLEncoded != resultStr) {
		t.Errorf(Base64URLEncoded + " not equal with " + resultStr)
	}
}

/*
	Test base64url.Decode function using a hex encoded string corresponding to an Base64URLEncoded string.
	First we decode the Base64URLEncoded string checking the error condition so we have a correctly encoded string.
	Then we hex encode the result and check it against the corresponding hex encoded string.
 */
func TestDecode(t *testing.T) {
	resultBytes, err := base64url.Decode(Base64URLEncoded)
	if(err != nil) {
		t.Errorf("Could not decode Base64URLEncoded string: ", err)
	}
	hexEncodedBytes := make([]byte, len(resultBytes) * 2)
	hex.Encode(hexEncodedBytes, resultBytes)
	resultStr := string(hexEncodedBytes)
	if(HEXEncoded != resultStr) {
		t.Errorf(HEXEncoded + " not equal with " + resultStr)		
	}
}