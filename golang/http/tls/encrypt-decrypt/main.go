package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

func main() {
	privKey, err := GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// privKey.PublicKey - contains the PublicKey

	text := []byte("Hello world")
	// The rand parameter is used as a source of entropy to ensure that
	// encrypting the same message twice doesn't result in the same
	// ciphertext.
	// PKCS = "Public Key Cryptography Standards"
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &privKey.PublicKey, text)
	if err != nil {
		log.Fatalf("failed to encrypt %v", err)
	}
	// the first parameter is the source of randomness and not required for decryption
	// the last parameter stores the options. If nil then PKCS is used for encryption
	originalText, err := privKey.Decrypt(nil, cipherText, nil)
	if err != nil {
		log.Fatalf("failed to decrypt just encrypted message: %v", err)
	}
	fmt.Printf("Original message: %s\n", string(text))
	fmt.Printf("Ciphertext %s\n", string(cipherText))
	fmt.Printf("Message after encryption/decryption: %s\n", string(originalText))
}

// GenerateKey generates private/public key pair
func GenerateKey() (*rsa.PrivateKey, error) {
	// 2048 - bit size of the key pair
	// rand.Reader is used as a source of random source
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	return privKey, err
}
