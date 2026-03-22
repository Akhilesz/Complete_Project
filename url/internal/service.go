package internal

import (
	"crypto/sha256"
	"encoding/binary"
)

const base62Alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateShortCode(longURL string) string {
	hash := sha256.Sum256([]byte(longURL))

	num := binary.BigEndian.Uint64(hash[:8])

	return encodeUsingBase62(num)
}

func encodeUsingBase62(num uint64) string {
	var shortCode string
	var mod uint64 = 62
	for num != 0 {
		var rem = num % mod
		shortCode = string(base62Alphabet[rem]) + shortCode
		num /= mod
	}
	return shortCode
}
