package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

// PKCS -  "Public Key Cryptography Standards", 15 is the last standard

func main() {
	msg := []byte("Hello there!")
	privKey, err := GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate a key pair: %v", err)
	}
	signature, err := GenerateSignature(privKey, msg)
	if err != nil {
		log.Fatalf("Failed to generate a signature: %v", err)
	}

	fmt.Println(VerifySignature(&privKey.PublicKey, []byte("Hello there!"), signature))              //true <nil>
	fmt.Println(VerifySignature(&privKey.PublicKey, []byte("Hello there"), signature))               //false crypto/rsa: verification error
	fmt.Println(VerifySignature(&privKey.PublicKey, []byte("Hello there!"), append(signature, 'c'))) //false crypto/rsa: verification error
}

func GenerateKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 4096)
}

func GenerateSignature(privKey *rsa.PrivateKey, msg []byte) ([]byte, error) {
	hash := sha256.Sum256(msg)
	return rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
}

func VerifySignature(pubKey *rsa.PublicKey, msg, signature []byte) (bool, error) {
	hash := sha256.Sum256(msg)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], signature)
	success := err == nil
	return success, err
}
