package pkg

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"
)

var salt = "How razorback-jumping frogs can level six piqued gymnasts!"
var secret []byte
var iv []byte

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
	return iv
}

func generateEncryptorStream(passphrase string) (cipher.Stream, error) {
	secret = generateKey(passphrase)
	encalg, err := aes.NewCipher(secret)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	iv = generateInitVector()
	stream := cipher.NewOFB(encalg, iv)
	return stream, nil
}

func Encrypt(passphrase string, from string, to string) error {
	log.Printf("Encrypt from: %s to %s passphrase %s\n", from, to, passphrase)

	ofile, err := os.Create(to)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	defer ofile.Close()

	ifile, err := os.Open(from)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	defer ifile.Close()

	secret = generateKey(passphrase)
	encalg, err := aes.NewCipher(secret)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	iv = generateInitVector()

	ofile.Write(secret)
	ofile.Write(iv)
	log.Printf("IV generated: %x\n", iv)

	str := cipher.NewCFBEncrypter(encalg, iv)

	idata, _ := ioutil.ReadAll(ifile)
	odata := make([]byte, len(idata))
	str.XORKeyStream(odata, idata)
	log.Printf("%s\n", string(idata))
	log.Printf("%x\n", odata)
	ofile.Write(odata)

	/*wtr := &cipher.StreamWriter{S: str, W: ofile}

	n, err := io.Copy(wtr, ifile)
	if err != nil {
		log.Printf("%s\n", err)
	}
	log.Printf("%d bytes copied\n", n)*/

	return nil
}

func Decrypt(passphrase string, from string, to string) error {
	log.Printf("Decrypt from: %s to %s passphrase %s\n", from, to, passphrase)

	ofile, err := os.Create(to)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	defer ofile.Close()

	ifile, err := os.Open(from)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	defer ifile.Close()

	secret = generateKey(passphrase)
	filepass := make([]byte, len(secret))

	bufr := bufio.NewReader(ifile)
	_, err = bufr.Read(filepass)

	//ifile.Read(filepass)
	log.Printf("Password:  %x\n", secret)
	log.Printf("From file: %x\n", filepass)

	res := bytes.Compare(secret, filepass)
	if res != 0 {
		log.Printf("Passphrase comparison failed\n")
		return nil
	}
	iv := make([]byte, aes.BlockSize)
	_, err = bufr.Read(iv)
	log.Printf("IV loaded: %x\n", iv)
	encalg, err := aes.NewCipher(filepass)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	st, _ := os.Stat(from)
	log.Printf("Size of input file %d\n", st.Size())
	buf := make([]byte, int(st.Size())-len(secret)-len(iv))
	//n, err := ifile.Read(buf)
	//log.Printf("%d bytes read. %s\n", n, err)
	bufr.Read(buf)
	log.Printf("Encrypted data:\n,%x\n", buf)
	obuf := make([]byte, len(buf))

	//str := cipher.NewOFB(encalg, iv)
	str := cipher.NewCFBDecrypter(encalg, iv)

	str.XORKeyStream(obuf, buf)

	n, err := ofile.Write(obuf)
	ofile.Close()
	log.Printf("%s\n", string(obuf))
	log.Printf("%x\n", obuf)
	log.Printf("%d bytes\n%s", n, err)

	/*

		rdr := &cipher.StreamReader{S: str, R: ifile}

		n, err := io.Copy(ofile, rdr)
		if err != nil {
			log.Printf("%s\n", err)
		}
		log.Printf("%d bytes copied\n", n)*/
	return nil
}
