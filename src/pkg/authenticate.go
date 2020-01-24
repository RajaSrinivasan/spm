package pkg

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
)

func LoadPublicKey(pubkeyfile string) (*rsa.PublicKey, error) {
	if Verbose {
		log.Printf("Loading public key %s\n", pubkeyfile)
	}
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
	if Verbose {
		log.Printf("Public key file %s parsed\n", pubkeyfile)
	}

	return pubkey, nil
}

func Authenticate(file string, sigfile string, pubkey *rsa.PublicKey) error {
	if pubkey == nil {
		log.Fatal("No public key provided")
	}

	hashed, _ := fileHash(file)
	sigbytes, _ := ioutil.ReadFile(sigfile)
	err := rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed, sigbytes)
	if err != nil {
		log.Printf("Verifying %s using signature: %s - %s\n", file, sigfile, err)
		return err
	}
	log.Printf("Authenticated the signature %s of file %s\n", sigfile, file)
	return nil
}

func AuthenticateFile(file string, sigfile string, pubkeyfile string) error {
	if Verbose {
		log.Printf("Authenticating %s signature %s publickey file %s\n", file, sigfile, pubkeyfile)
	}
	rsapubkey, err := LoadPublicKey(pubkeyfile)
	if err != nil {
		return err
	}
	err = Authenticate(file, sigfile, rsapubkey)
	return err
}

func AuthenticateFiles(files []string, pub string) error {
	if Verbose {
		log.Printf("Authenticating %d files using %s\n", len(files), pub)
	}

	pubkeyfile, err := LoadPublicKey(pub)
	if err != nil {
		log.Printf("Cannot authenticate files. Unable to load Public Key file\n")
		return err
	}
	for _, f := range files {
		sigfile := f + signatureFileType
		err = Authenticate(f, sigfile, pubkeyfile)
		if err != nil {
			log.Printf("Autentication failed for %s\n", f)
		}
	}
	return nil
}
