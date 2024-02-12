package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func LoadPrivateKey() *rsa.PrivateKey {
	filePath := fmt.Sprintf("%s/key/pkcs8_privateKey.pem", os.Getenv("APP_PATH"))
	privKey, err := os.ReadFile(filePath)

	dec, _ := pem.Decode(privKey)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rsaPriv, err := x509.ParsePKCS8PrivateKey(dec.Bytes)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return rsaPriv.(*rsa.PrivateKey)
}

func LoadPublicKey() *rsa.PublicKey {
	filePath := fmt.Sprintf("%s/key/pkcs8.pub", os.Getenv("APP_PATH"))

	pubKey, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	decpub, _ := pem.Decode(pubKey)
	rsapubd, err := x509.ParsePKIXPublicKey(decpub.Bytes)
	s := rsapubd.(*rsa.PublicKey)
	if err != nil {
		return nil
	}
	return s
}
