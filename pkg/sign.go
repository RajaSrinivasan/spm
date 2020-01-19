package pkg

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

var Verbose = true

const rsaKeySize = 2048
const signatureFileType = ".sig"
const DefaultPrivateKeyFileName = "private.pem"
const DefaultPublicKeyFileName = "public.pem"

func showPrivateKey(pvt *rsa.PrivateKey) {
	log.Printf("%s\n", pvt.D.Text(16))
	pvtkeybytes, _ := x509.MarshalPKCS8PrivateKey(pvt)

	log.Printf("PvtKeyBytes: %x\n", pvtkeybytes)

	pub := pvt.Public()
	rsapub := pub.(*rsa.PublicKey)

	log.Printf("Public Key Size %d, Exponent %d\n", rsapub.Size(), rsapub.E)
	log.Printf("PublicKey: %s\n", rsapub.N.Text(16))
}

func LoadPrivateKey(keyfile string) (*rsa.PrivateKey, error) {
	if Verbose {
		log.Printf("Loading private key %s\n", keyfile)
	}
	keybytes, _ := ioutil.ReadFile(keyfile)
	key, err := ssh.ParseRawPrivateKey(keybytes)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}

	return key.(*rsa.PrivateKey), nil
}

func loadPrivateKeyWithPassphrase(keyfile string, passphrase string) (*rsa.PrivateKey, error) {
	log.Printf("Passphrase %s\n", passphrase)
	keybytes, _ := ioutil.ReadFile(keyfile)

	block, _ := pem.Decode(keybytes)
	if block == nil {
		log.Printf("Unable to decode. block is nil\n")
		return nil, nil
	}
	if block.Type != "PRIVATE KEY" {
		log.Printf("Wrong block Type %s\n", block.Type)
		return nil, nil
	}

	rawkey, err := x509.DecryptPEMBlock(block, []byte(passphrase))
	if err != nil {
		log.Printf("Error decrypting PEM block")
		return nil, err
	}
	log.Printf("%q\n", rawkey)

	key, err := x509.ParsePKCS8PrivateKey(rawkey)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	if Verbose {
		log.Printf("%q\n", key)
	}
	return key.(*rsa.PrivateKey), nil
}

func GenerateKeys(priv, pub string) error {
	if Verbose {
		log.Printf("Generating keys Private: %s and Public: %s\n", priv, pub)
	}
	privkey, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	pvtder, err := x509.MarshalPKCS8PrivateKey(privkey)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	//log.Printf("Converted to DER key\n")

	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pvtder}
	privpem, err := os.Create(priv)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	defer privpem.Close()
	err = pem.Encode(privpem, block)
	if err != nil {
		log.Printf("%s unable to PEM Encode\n", err)
		return err
	}
	var pubkey rsa.PublicKey
	pubkey = privkey.PublicKey

	pubder := x509.MarshalPKCS1PublicKey(&pubkey)

	pblock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubder}

	pubpem, err := os.Create(pub)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	defer pubpem.Close()
	err = pem.Encode(pubpem, pblock)
	if err != nil {
		log.Printf("%s unable to PEM Encode\n", err)
		return err
	}

	return nil
}

func generateKeysWithPassphrase(priv, pub string, passphrase string) error {
	privkey, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	pvtder, err := x509.MarshalPKCS8PrivateKey(privkey)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	log.Printf("Converted to DER key\n")
	block, err := x509.EncryptPEMBlock(rand.Reader, "PRIVATE KEY", pvtder, []byte(passphrase), x509.PEMCipherAES256)
	if err != nil {
		log.Printf("%s", err)
		return nil
	}
	privpem, err := os.Create(priv)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	defer privpem.Close()

	err = pem.Encode(privpem, block)
	if err != nil {
		log.Printf("%s unable to PEM Encode\n", err)
		return err
	}
	var pubkey rsa.PublicKey
	pubkey = privkey.PublicKey

	pubder := x509.MarshalPKCS1PublicKey(&pubkey)

	pblock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubder}

	pubpem, err := os.Create(pub)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	defer pubpem.Close()
	err = pem.Encode(pubpem, pblock)
	if err != nil {
		log.Printf("%s unable to PEM Encode\n", err)
		return err
	}

	return nil
}

func fileHash(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, nil
	}
	defer f.Close()

	h := sha256.New()
	io.Copy(h, f)
	hash := h.Sum(nil)
	return hash, nil
}

func Sign(file string, sigfile string, pvt *rsa.PrivateKey) error {

	log.Printf("Signing %s creating %s\n", file, sigfile)
	datahash, _ := fileHash(file)
	if Verbose {
		log.Printf("Datahash: %x\n", datahash)
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, pvt, crypto.SHA256, datahash[:])
	if err != nil {
		log.Printf("Error from signing: %s\n", err)
		return err
	}

	sigf, _ := os.Create(sigfile)
	defer sigf.Close()
	sigf.Write(signature)
	if Verbose {
		fmt.Printf("Signature: %x\n", signature)
	}
	return nil
}

func signFileWithPassphrase(file string, sigfile string, pvtkeyfile string, passphrase string) error {
	var err error
	var rsapvtkey *rsa.PrivateKey
	if Verbose {
		log.Printf("Signing %s with %s to generate %s\n", file, pvtkeyfile, sigfile)
	}
	if len(passphrase) > 0 {
		rsapvtkey, err = loadPrivateKeyWithPassphrase(pvtkeyfile, passphrase)
	} else {
		rsapvtkey, err = LoadPrivateKey(pvtkeyfile)
	}
	if err != nil {
		return err
	}
	err = Sign(file, sigfile, rsapvtkey)
	return err
}

func SignFile(file string, sigfile string, pvtkeyfile string) error {
	var err error
	var rsapvtkey *rsa.PrivateKey
	if Verbose {
		log.Printf("Signing %s with %s to generate %s\n", file, pvtkeyfile, sigfile)
	}

	rsapvtkey, err = LoadPrivateKey(pvtkeyfile)
	if err != nil {
		return err
	}
	err = Sign(file, sigfile, rsapvtkey)
	return err
}

func SignFiles(files []string, pvtkeyfile string) error {
	log.Printf("Signing using %s of %d files\n", pvtkeyfile, len(files))
	rsapvtkey, err := LoadPrivateKey(pvtkeyfile)
	if err != nil {
		return err
	}
	for _, f := range files {
		sigfname := f + signatureFileType
		err = Sign(f, sigfname, rsapvtkey)
		if err != nil {
			return err
		}
	}
	return nil
}
