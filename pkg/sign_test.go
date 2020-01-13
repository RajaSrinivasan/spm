package pkg

import (
	"log"
	"testing"
)

func TestGenerateKey(t *testing.T) {

	_, err := generatePrivateKey()
	if err == nil {
		log.Printf("Generated\n")
	}

	_, err = generatePrivateKey()
	if err == nil {
		log.Printf("Generated\n")
	}
}

func TestSign(t *testing.T) {
	pvt, _ := generatePrivateKey()
	sign("sign.go", "sign.go.sig", pvt)
	sign("sign.go", "sign.go.2.sig", pvt)
	sign("sign_test.go", "sign_test.go.sig", pvt)
	sign("sign_test.go", "sign_test.go.2.sig", pvt)
}

func TestGenerate(t *testing.T) {
	GenerateKeyPair("privatekey", "publickey")
	GenerateKeyPair("privatekey2", "publickey2")
}

func TestSignExternalKey(t *testing.T) {
	Sign("sign.go", "sign.go.sig", "privatekey")
	Sign("sign.go", "sign.go.2.sig", "privatekey2")
	Sign("sign_test.go", "sign_test.go.sig", "privatekey")
	Sign("sign_test.go", "sign_test.go.2.sig", "privatekey2")
}

func TestVerifyExternalKey(t *testing.T) {
	Verify("sign.go", "sign.go.sig", "publickey")
	Verify("sign.go", "sign.go.2.sig", "publickey2")
	Verify("sign_test.go", "sign_test.go.sig", "publickey")
	Verify("sign_test.go", "sign_test.go.2.sig", "publickey2")
}
