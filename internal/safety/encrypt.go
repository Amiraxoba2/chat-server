package safety

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func Encrypt(msg string, key rsa.PublicKey) string {
	label := []byte("OAEP Encrypted")
	rng := rand.Reader
	ciphertext, _ := rsa.EncryptOAEP(sha256.New(), rng, &key, []byte(msg), label)
	return base64.StdEncoding.EncodeToString(ciphertext)
}
