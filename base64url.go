package base64url

/*
	Implementation of base64url removing the traling = and also changing +,/ to the more
	url friendly -,_ characters.

	The decode implementation is reversing the process before running DecodeString from the base64 package.
*/

import (
	"encoding/base64"
	"math/rand"
	"strings"
)

const PossibleValues string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func Encode(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}

func Decode(str string) ([]byte, error) {
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	for len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}

// generates n number of valid base64 URL characters
func Rand(size int) string {
	var buff string
	for ; size > 0; size-- {
		random := rand.Intn(len(PossibleValues))
		buff = buff + PossibleValues[random:random+1]
	}
	return buff
}
