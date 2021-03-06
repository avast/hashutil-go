// Code generated by github.com/jasei/hashutil/generator DO NOT EDIT
package hashutil

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"strings"
)

// Md5 type represents Md5 checksum
type Md5 []byte

// StringToMd5 return a new Md5 checksum from string (hex) representation
func StringToMd5(hexString string) (Md5, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return Md5{}, err
	}

	return BytesToMd5(bytes)
}

// BytesToMd5 return a new Md5 checksum from bytes (binary) representation
func BytesToMd5(bytes []byte) (Md5, error) {
	if len(bytes) != 16 {
		return Md5{}, fmt.Errorf("Hash function Md5 must have a length of 16 bytes (actual have %d)", len(bytes))
	}

	return Md5(bytes), nil

}

//HashToMd5 return a new Md5 checksum from hash.Hash representation
// HashToMd5 convert hashutil.Hash to Md5
func HashToMd5(h hash.Hash) (Md5, error) {
	return BytesToMd5(h.Sum(nil))
}

// EmptyMd5 return Md5 of empty file
func EmptyMd5() Md5 {
	h, _ := StringToMd5("d41d8cd98f00b204e9800998ecf8427e")
	return h
}

// Equal return true if is Md5s equal
func (h Md5) Equal(o Md5) bool {
	return bytes.Equal(h, o)
}

// String return (hex) string representation of Md5
func (h Md5) String() string {
	return hex.EncodeToString(h)
}

// UpperString return (hex) string representation in upper case of Md5
func (h Md5) UpperString() string {
	return strings.ToUpper(hex.EncodeToString(h))
}

// ToBytes return []byte of hashutil.Md5
func (h Md5) ToBytes() []byte {
	return h
}

// ToBase64 return base64 representation of Md5
func (h Md5) ToBase64() string {
	return base64.StdEncoding.EncodeToString(([]byte)(h))
}

// IsEmpty return true if is Md5 'empty hash'
func (h Md5) IsEmpty() bool {
	return EmptyMd5().Equal(h)
}
