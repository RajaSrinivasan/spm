package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const salt = "How razorback-jumping frogs can level six piqued gymnasts!"

func generateKey(passphrase string) []byte {
	salted := passphrase + salt
	hash := sha256.New()
	hash.Write([]byte(salted))
	return hash.Sum(nil)
}

func generateInitVector() []byte {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		log.Printf("%s\n", err)
	}
	if Verbose {
		log.Printf("IV: %x\n", iv)
	}
	return iv
}

func Encrypt(passphrase string, from string, to string) error {
	log.Printf("Encrypt from: %s to %s passphrase %s\n", from, to, passphrase)

	ofile, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer ofile.Close()

	ifile, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}
	defer ifile.Close()

	secret := generateKey(passphrase)
	encalg, err := aes.NewCipher(secret)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	iv := generateInitVector()

	ofile.Write(secret)
	ofile.Write(iv)

	str := cipher.NewOFB(encalg, iv)

	idata, _ := ioutil.ReadAll(ifile)
	odata := make([]byte, len(idata))
	str.XORKeyStream(odata, idata)

	ofile.Write(odata)

	return nil
}

func Decrypt(passphrase string, from string, to string) error {
	log.Printf("Decrypt from: %s to %s passphrase %s\n", from, to, passphrase)

	ofile, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer ofile.Close()

	ifile, err := os.Open(from)
	if err != nil {
		log.Fatal(err)
	}
	defer ifile.Close()

	secret := generateKey(passphrase)
	filepass := make([]byte, len(secret))
	bytesin, err := io.ReadFull(ifile, filepass)
	if err != nil {
		log.Fatal(err)
	}
	if Verbose {
		log.Printf("%d bytes read for password\n", bytesin)
	}
	res := bytes.Compare(secret, filepass)
	if res != 0 {
		log.Printf("Passphrase comparison failed\n")
		return errors.New("Passphrase comparison failed")
	}

	iv := make([]byte, aes.BlockSize)
	bytesin, err = io.ReadFull(ifile, iv)
	if err != nil {
		log.Fatal(err)
	}
	if Verbose {
		log.Printf("%d bytes read for IV\n", bytesin)
		log.Printf("IV: %x\n", iv)
	}
	encalg, err := aes.NewCipher(filepass)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	st, _ := os.Stat(from)
	buf := make([]byte, int(st.Size())-len(secret)-len(iv))
	nbytesin, err := io.ReadFull(ifile, buf)
	if err != nil {
		log.Fatal(err)
	}
	if Verbose {
		log.Printf("%d bytes read\n", nbytesin)
	}
	obuf := make([]byte, len(buf))

	str := cipher.NewOFB(encalg, iv)
	str.XORKeyStream(obuf, buf)

	nbytesout, err := ofile.Write(obuf)
	if err != nil {
		log.Fatal(err)
	}
	if Verbose {
		log.Printf("%d bytes written\n", nbytesout)
	}
	ofile.Close()
	return nil
}
