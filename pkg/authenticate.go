package pkg

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func loadPublicKey(pubkeyfile string) (*rsa.PublicKey, error) {

	keybytes, err := ioutil.ReadFile(pubkeyfile)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, err
	}

	block, _ := pem.Decode(keybytes)
	if block == nil {
		log.Printf("Unable to decode. block is nil\n")
		return nil, nil
	}
	if block.Type != "PUBLIC KEY" {
		log.Printf("Wrong block Type %s\n", block.Type)
		return nil, nil
	}

	pubkey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	log.Printf("Public key parsed\n")

	/*decode string ssh-rsa format to native type
	pub_key, err := ssh.DecodePublicKey(string(pubbytes))
	if err != nil {
		log.Printf("%v\n", err)
	}

	rsapubkey := pub_key.(*rsa.PublicKey)
	return rsapubkey, nil*/
	return pubkey, nil
}

func authenticate(file string, sigfile string, pubkey *rsa.PublicKey) error {
	databytes, _ := ioutil.ReadFile(file)
	h := sha256.New()
	h.Write(databytes)
	hashed := h.Sum(nil)

	sigbytes, _ := ioutil.ReadFile(sigfile)
	err := rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed, sigbytes)
	if err != nil {
		log.Printf("Verifying %s using signature: %s - %s\n", file, sigfile, err)
		return err
	}
	log.Printf("Verified the signature %s of file %s\n", sigfile, file)
	return nil
}

func Authenticate(file string, sigfile string, pubkeyfile string) error {

	rsapubkey, err := loadPublicKey(pubkeyfile)
	if err != nil {
		return err
	}
	log.Printf("Loaded public key %s\n", pubkeyfile)
	err = authenticate(file, sigfile, rsapubkey)
	return err
}

func AuthenticateFiles(files []string, pub string) error {
	log.Printf("Authenticating using %s of %d files\n", pub, len(files))
	pubkeyfile, err := loadPublicKey(pub)
	if err != nil {
		log.Printf("Cannot authenticate files. Unable to load Public Key file\n")
		return err
	}
	for _, f := range files {
		sigfile := f + signatureFileType
		err = authenticate(f, sigfile, pubkeyfile)
		if err != nil {
			return err
		}
	}
	return nil
}
