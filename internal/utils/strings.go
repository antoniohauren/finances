package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func bytesToHex(b []byte) string {
	return strings.ToUpper(hex.EncodeToString(b))
}

func GenerateUserCode() string {
	part := make([]byte, 4)

	rand.Read(part)

	return fmt.Sprintf("%s-%s", bytesToHex(part[:2]), bytesToHex(part[2:]))
}
