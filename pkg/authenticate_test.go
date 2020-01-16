package pkg

import (
	"log"
	"testing"
)

var pubKeyFile = "/Users/rajasrinivasan/.ssh/id_rsa.pub"
var pubKeyPEM = "public.pem"

func TestAuthenticate(t *testing.T) {
	//t.Println("Testing authentication of digital signatures")
	Authenticate("sign.go", "sign.go.sig", pubKeyPEM)
	Authenticate("sign_test.go", "sign_test.go.sig", pubKeyPEM)
}

func TestAuthenticateFiles(t *testing.T) {
	files := []string{"sign.go", "sign_test.go"}
	AuthenticateFiles(files, pubKeyFile)
	files = []string{"authenticate.go", "authenticate_test.go", "sign.go", "sign_test.go"}
	AuthenticateFiles(files, pubKeyFile)
}

func TestLoadPrublicKey(t *testing.T) {
	pubkey, _ := loadPublicKey(pubKeyPEM)
	log.Printf("%v\n", pubkey)
}
