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

// Sha256 type represents Sha256 checksum
type Sha256 []byte

// StringToSha256 return a new Sha256 checksum from string (hex) representation
func StringToSha256(hexString string) (Sha256, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return Sha256{}, err
	}

	return BytesToSha256(bytes)
}

// BytesToSha256 return a new Sha256 checksum from bytes (binary) representation
func BytesToSha256(bytes []byte) (Sha256, error) {
	if len(bytes) != 32 {
		return Sha256{}, fmt.Errorf("Hash function Sha256 must have a length of 32 bytes (actual have %d)", len(bytes))
	}

	return Sha256(bytes), nil

}

//HashToSha256 return a new Sha256 checksum from hash.Hash representation
// HashToSha256 convert hashutil.Hash to Sha256
func HashToSha256(h hash.Hash) (Sha256, error) {
	return BytesToSha256(h.Sum(nil))
}

// EmptySha256 return Sha256 of empty file
func EmptySha256() Sha256 {
	h, _ := StringToSha256("e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855")
	return h
}

// Equal return true if is Sha256s equal
func (h Sha256) Equal(o Sha256) bool {
	return bytes.Equal(h, o)
}

// String return (hex) string representation of Sha256
func (h Sha256) String() string {
	return hex.EncodeToString(h)
}

// UpperString return (hex) string representation in upper case of Sha256
func (h Sha256) UpperString() string {
	return strings.ToUpper(hex.EncodeToString(h))
}

// ToBytes return []byte of hashutil.Sha256
func (h Sha256) ToBytes() []byte {
	return h
}

// ToBase64 return base64 representation of Sha256
func (h Sha256) ToBase64() string {
	return base64.StdEncoding.EncodeToString(([]byte)(h))
}

// IsEmpty return true if is Sha256 'empty hash'
func (h Sha256) IsEmpty() bool {
	return EmptySha256().Equal(h)
}
