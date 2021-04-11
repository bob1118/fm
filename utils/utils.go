package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

//string compare case-insensitivity
func IsEqual(s string, d string) (b bool) {
	return strings.EqualFold(s, d)
}

//MakeA1Hash function.
func MakeA1Hash(in string) (s string) {
	h := md5.New()
	io.WriteString(h, in)
	return fmt.Sprintf("%x", h.Sum(nil))
}
