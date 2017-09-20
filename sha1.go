package hashutil

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Sha1 []byte

var badSha1Length = fmt.Errorf("Sha1 must have a length of %d bytes", sha1.Size)

// convert type Sha1 to string representation (lower case)
func (hash Sha1) String() string {
	return hex.EncodeToString(hash)
}

// StringToSha1 validate and convert a string to Sha1 type
func StringToSha1(str string) (Sha1, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}

	if len(bytes) != sha1.Size {
		return nil, badSha1Length
	}

	return bytes, nil
}

// BytesToSha1 validate and convert []bytes to Sha1 type
func BytesToSha1(bytes []byte) (Sha1, error) {
	if len(bytes) != sha1.Size {
		return nil, badSha1Length
	}

	return bytes, nil
}

// EmptySha1 return type Sha1 of empty (zero-length) string
func EmptySha1() Sha1 {
	return sha1.New().Sum(nil)
}