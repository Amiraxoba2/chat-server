package safety

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func LoadKey() *rsa.PrivateKey {
	pemBytes, _ := os.ReadFile("resource/key.rsa")
	block, _ := pem.Decode(pemBytes)
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey
}
