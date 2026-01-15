package http

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func signHMAC(secret string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return hex.EncodeToString(mac.Sum(nil))
}

func verifyHMAC(secret string, payload []byte, signature string) bool {
	expected := signHMAC(secret, payload)
	return hmac.Equal([]byte(expected), []byte(signature))
}
