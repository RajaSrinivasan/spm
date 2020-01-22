package pkg

import (
	"log"
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
	Encrypt("Thisisabadpassph", "crypt_test.go", "crypt_test_go.enc")
	Encrypt("Thisisabadpassphrase", "crypt.go", "crypt_go.enc")
}

func TestEncryptFileBig(t *testing.T) {
	Encrypt("Thisisabadpassph", "/Users/rajasrinivasan/Prj/work/sp_clear.spm", "../tests/sp.spm")
}

func TestDecryptFile(t *testing.T) {
	Decrypt("Thisisabadpassphrase", "crypt_test_go.enc", "crypt_test_go")
	Decrypt("Thisisabadpassph", "crypt_test_go.enc", "crypt_test_go")
	Decrypt("Thisisabadpassphrase", "crypt_go.enc", "crypt_go")
}

func TestDecryptFileBig(t *testing.T) {
	Decrypt("Thisisabadpassph", "../tests/sp.spm", "../tests/sp_clear.spm")
}
