package main

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GenerateEightCharID() string {
	// Need at least 6 bytes to get 8 base64 URL-safe characters (6*4/3 = 8 chars)
	bytes, _ := GenerateRandomBytes(6)

	// Base64 URL encoding ensures the string is URL-safe and compact
	id := base64.RawURLEncoding.EncodeToString(bytes)
	// Truncate to ensure exactly 8 characters if the encoding adds padding (RawURLEncoding doesn't)
	// The length will naturally be 8 characters.
	return id
}
