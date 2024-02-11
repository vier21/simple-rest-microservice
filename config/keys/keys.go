package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"
)

func LoadPrivatKey() (*rsa.PrivateKey, error) {
	filePath := fmt.Sprintf(os.Getenv("APP_PATH"), "/key/privateKey.pem")
	byteKey, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	rsa, err := x509.ParsePKCS1PrivateKey(byteKey)

	if err != nil {
		return nil, err
	}

	return rsa, nil
}

func LoadPublicKey() (*rsa.PublicKey, error) {
	filePath := fmt.Sprintf(os.Getenv("APP_PATH"), "/key/publicKey.pub")
	byteKey, err := os.ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	rsa, err := x509.ParsePKCS1PublicKey(byteKey)

	if err != nil {
		return nil, err
	}

	return rsa, nil
}
