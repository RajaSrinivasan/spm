package pkg

import (
	"crypto/cipher"
	"io"
	"log"
	"os"
	"testing"
)

func TestSalting(t *testing.T) {
	hash := generateKey("This is a bad passphrase")
	log.Printf("Hashed passphrase : length %d, %x\n", len(hash), hash)
	hash = generateKey("This is a bad passphrase")
	log.Printf("Hashed passphrase : length %d, %x\n", len(hash), hash)
}

func TestIV(t *testing.T) {
	iv := generateInitVector()
	log.Printf("Init Vector: length %d %x\n", len(iv), iv)
	iv = generateInitVector()
	log.Printf("Init Vector: length %d %x\n", len(iv), iv)
	iv = generateInitVector()
	log.Printf("Init Vector: length %d %x\n", len(iv), iv)
}

func TestGenerateEncryptorStream(t *testing.T) {
	str, err := generateEncryptorStream("This is also a bad passphrase")
	if err != nil {
		log.Printf("Generating Encryptor stream failed\n")
	}
	ofile, _ := os.Create("encrypted")
	ifile, _ := os.Open("crypt_test.go")
	wtr := &cipher.StreamWriter{S: str, W: ofile}
	io.Copy(wtr, ifile)

}

func TestEncryptFile(t *testing.T) {
	Encrypt("Thisisabadpassph", "crypt_test.go", "crypt_test_go.enc")
}

func TestDecryptFile(t *testing.T) {
	Decrypt("Thisisabadpassph", "crypt_test_go.enc", "crypt_test_go")
}
