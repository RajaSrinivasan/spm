package pkg

import (
	"log"
	"path/filepath"
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

func TestEncryptFile(t *testing.T) {
	Encrypt("Thisisabadpassph", "crypt_test.go", "/tmp/crypt_test_go.enc")
	Encrypt("Thisisabadpassphrase", "crypt.go", "/tmp/crypt_go.enc")
}

func TestEncryptFileBig(t *testing.T) {
	TestPackfilesBig(t)
	Encrypt("Thisisabadpassph", filepath.Join(WorkDir, "bigpack.tgz"), filepath.Join(WorkDir, "bigpack.spm"))
}

func TestDecryptFile(t *testing.T) {
	TestEncryptFile(t)
	Decrypt("Thisisabadpassphrase", "/tmp/crypt_test_go.enc", "/tmp/crypt_test_go")
	Decrypt("Thisisabadpassph", "/tmp/crypt_test_go.enc", "/tmp/crypt_test_go")
	Decrypt("Thisisabadpassphrase", "/tmp/crypt_go.enc", "/tmp/crypt_go")
}

func TestDecryptFileBig(t *testing.T) {
	TestEncryptFileBig(t)
	Decrypt("Thisisabadpassph", filepath.Join(WorkDir, "bigpack.spm"), "/tmp/bigpack.tgz")
}
