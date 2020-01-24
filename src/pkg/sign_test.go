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
	privkey, err := LoadPrivateKey("../tests/private.pem")
	if err == nil {
		showPrivateKey(privkey)
	}
	privkey, _ = LoadPrivateKey("../tests/private_missing.pem")

}

func TestLoadPrivateKeyfileWithPassphrase(t *testing.T) {
	t.Log("Testing Loading private pem files")
	privkey, err := loadPrivateKeyWithPassphrase("pwd_private.pem", goodpassphrase)
	if err == nil {
		showPrivateKey(privkey)
	}
}

func TestSignFile(t *testing.T) {
	err := SignFile("sign.go", "../tests/sign.go.sig", "../tests/private.pem")
	if err == nil {
		log.Printf("Signed\n")
	}
	err = SignFile("sign_test.go", "../tests/sign_test.go.sig", "../tests/private.pem")
	if err == nil {
		log.Printf("Signed\n")
	}
	SignFile("sign_test.go", "../tests/sign_test1.go.sig", "../tests/bad_private.pem")
	SignFile("sign_test.go", "../tests/sign_test2.go.sig", "../tests/bad_private.pem")
}

func TestSign(t *testing.T) {
	keyfile, err := LoadPrivateKey("../tests/private.pem")
	if err != nil {
		log.Panic(err)
	}
	err = Sign("sign.go", "../tests/sign.go.sig", keyfile)
	if err == nil {
		log.Printf("Signed\n")
	}
	err = Sign("sign_test.go", "../tests/sign_test.go.sig", keyfile)
	if err == nil {
		log.Printf("Signed\n")
	}
	err = Sign("sign_test.go", "../tests/sign_test.go.sig", nil)
	if err == nil {
		log.Printf("Signed\n")
	}
}

func TestSignWithPassphrase(t *testing.T) {
	err := signFileWithPassphrase("sign.go", "sign.go.sig", "pwd_private.pem", goodpassphrase)
	if err == nil {
		log.Printf("Signed\n")
	}
}
