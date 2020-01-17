package pkg

import (
	"log"
	"os"
	"testing"
)

var pubKeyFile = "/Users/rajasrinivasan/.ssh/id_rsa.pub"
var pubKeyPEM = "../tests/public.pem"

func TestLoadPrublicKey(t *testing.T) {
	pubkey, _ := LoadPublicKey(pubKeyPEM)
	log.Printf("%v\n", pubkey)
	pubkey, _ = LoadPublicKey("public.pem")
	pubkey, _ = LoadPublicKey("../tests/badpublic.pem")

}

func TestAuthenticateFile(t *testing.T) {
	//t.Println("Testing authentication of digital signatures")
	AuthenticateFile("sign.go", "../tests/sign.go.sig", pubKeyPEM)
	AuthenticateFile("sign_test.go", "../tests/sign_test.go.sig", pubKeyPEM)
}

func TestAuthenticate(t *testing.T) {
	pubkey, _ := LoadPublicKey(pubKeyPEM)
	Authenticate("sign.go", "../tests/sign.go.sig", pubkey)
	Authenticate("sign_test.go", "../tests/sign_test.go.sig", pubkey)
}

func TestAuthenticateFiles(t *testing.T) {
	os.Chdir("../tests")
	files := []string{"lsfiles.txt", "private.pem", "public.pem"}
	SignFile("lsfiles.txt", "lsfiles.txt.sig", "private.pem", "")
	SignFile("private.pem", "private.pem.sig", "private.pem", "")
	AuthenticateFiles(files, pubKeyPEM)

}
