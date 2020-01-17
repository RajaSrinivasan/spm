package pkg

import (
	"log"
	"testing"
)

var goodpassphrase = "Thisisagoodpassphrase"

func TestGenerateKeys(t *testing.T) {
	t.Log("Testing Key generation")
	err := GenerateKeys("../tests/private.pem", "../tests/public.pem")
	if err == nil {
		log.Printf("Generated\n")
	}
}

func TestGenerateKeysWithPassphrase(t *testing.T) {
	t.Log("Testing Key generation")
	err := generateKeysWithPassphrase("pwd_private.pem", "pwd_public.pem", goodpassphrase)
	if err == nil {
		log.Printf("Generated\n")
	}
}

func TestLoadPrivateKeyfile(t *testing.T) {
	t.Log("Testing Loading private pem files")
	privkey, err := LoadPrivateKey("private.pem")
	if err == nil {
		showPrivateKey(privkey)
	}
}

func TestLoadPrivateKeyfileWithPassphrase(t *testing.T) {
	t.Log("Testing Loading private pem files")
	privkey, err := loadPrivateKeyWithPassphrase("pwd_private.pem", goodpassphrase)
	if err == nil {
		showPrivateKey(privkey)
	}
}

func TestSign(t *testing.T) {
	err := Sign("sign.go", "sign.go.sig", "private.pem", "")
	if err == nil {
		log.Printf("Signed\n")
	}
	err = Sign("sign_test.go", "sign_test.go.sig", "private.pem", "")
	if err == nil {
		log.Printf("Signed\n")
	}
}

func TestSignWithPassphrase(t *testing.T) {
	err := Sign("sign.go", "sign.go.sig", "pwd_private.pem", goodpassphrase)
	if err == nil {
		log.Printf("Signed\n")
	}
}
