package safety

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func GenerateKey() {
	privKey, _ := rsa.GenerateKey(rand.Reader, 4096)
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privKey),
		},
	)
	os.WriteFile("resource/key.rsa", keyPEM, 0700)
}
