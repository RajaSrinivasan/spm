package pkg

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"fmt"
	"os"

	"io/ioutil"
	"log"
)

const rsaKeySize = 2048

func showPrivateKey(pvt *rsa.PrivateKey) {
	log.Printf("%s\n", pvt.D.Text(16))
	pvtkeybytes, _ := x509.MarshalPKCS8PrivateKey(pvt)

	log.Printf("PvtKeyBytes: %x\n", pvtkeybytes)

	pub := pvt.Public()
	rsapub := pub.(*rsa.PublicKey)

	log.Printf("Public Key Size %d, Exponent %d\n", rsapub.Size(), rsapub.E)
	log.Printf("PublicKey: %s\n", rsapub.N.Text(16))
}

func generatePrivateKey() (*rsa.PrivateKey, error) {
	pvt, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		log.Printf("%s\n", err)
		return nil, err
	}
	showPrivateKey(pvt)
	return pvt, nil
}

func GenerateKeyPair(priv, pub string) error {
	pvt, err := rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	pvtkeybytes, err := x509.MarshalPKCS8PrivateKey(pvt)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	privf, _ := os.Create(priv)
	defer privf.Close()
	privf.Write(pvtkeybytes)

	pubkey := pvt.Public()
	rsapub := pubkey.(*rsa.PublicKey)
	pubkeybytes, err := x509.MarshalPKIXPublicKey(rsapub)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	pubf, _ := os.Create(pub)
	defer pubf.Close()
	pubf.Write(pubkeybytes)

	return nil
}

func sign(file string, sigfile string, pvt *rsa.PrivateKey) error {

	log.Printf("Signing %s creating %s\n", file, sigfile)

	databytes, _ := ioutil.ReadFile(file)
	h := sha256.New()
	h.Write(databytes)
	datahash := h.Sum(nil)
	log.Printf("Datahash: %x\n", datahash)

	signature, err := rsa.SignPKCS1v15(rand.Reader, pvt, crypto.SHA256, datahash[:])
	if err != nil {
		log.Printf("Error from signing: %s\n", err)
		return err
	}

	sigf, _ := os.Create(sigfile)
	defer sigf.Close()
	sigf.Write(signature)

	fmt.Printf("Signature: %x\n", signature)
	return nil
}

func Sign(file string, sigfile string, pvtkeyfile string) error {
	pvtbytes, err := ioutil.ReadFile(pvtkeyfile)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	pvtkey, err := x509.ParsePKCS8PrivateKey(pvtbytes)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	rsapvtkey := pvtkey.(*rsa.PrivateKey)
	err = sign(file, sigfile, rsapvtkey)
	return err
}

func verify(file string, sigfile string, pubkey *rsa.PublicKey) error {
	databytes, _ := ioutil.ReadFile(file)
	h := sha256.New()
	h.Write(databytes)
	hashed := h.Sum(nil)

	sigbytes, _ := ioutil.ReadFile(sigfile)
	err := rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed, sigbytes)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	return nil
}

func Verify(file string, sigfile string, pubkeyfile string) error {
	log.Printf("Verifying %s signature %s using %s\n", file, sigfile, pubkeyfile)
	pubbytes, err := ioutil.ReadFile(pubkeyfile)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}
	log.Printf("Loaded %s %d bytes\n", pubkeyfile, len(pubbytes))
	pubkey, err := x509.ParsePKIXPublicKey(pubbytes)
	if err != nil {
		log.Printf("%s\n", err)
		return err
	}

	rsapubkey := pubkey.(*rsa.PublicKey)
	err = verify(file, sigfile, rsapubkey)

	return nil
}
