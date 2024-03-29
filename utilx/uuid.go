/*
From: https://github.com/dailymotion/allure-go/blob/c1bc79874bb49b81e0fd83e7bdf7a88ef36cb6b5/random.go
*/
package utilx

import (
	"crypto/rand"
	"fmt"
)

// GenerateUUID UUID格式: 2031ec7b-2506-4248-aeb2-b1d2209ccbb9
func GenerateUUID() string {
	var entropy = make([]byte, 16)
	rand.Read(entropy)
	entropy[6] = (entropy[6] & 0x0f) | 0x40
	entropy[8] = (entropy[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", entropy[0:4], entropy[4:6], entropy[6:8], entropy[8:10], entropy[10:16])
}
