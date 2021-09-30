package main

import (
	"log"
	"os"
)

// The Ed25519 was introduced on OpenSSH version 6.5.
// It’s the EdDSA implementation using the Twisted Edwards curve.
// It’s using elliptic curve cryptography that offers a better security with faster performance compared to DSA or ECDSA.
// Today, the RSA is the most widely used public-key algorithm for SSH key.
// But compared to Ed25519, it’s slower and even considered not safe if it’s generated with the key smaller than 2048-bit length.
// The Ed25519 public-key is compact. It only contains 68 characters, compared to RSA 3072 that has 544 characters.
// Generating the key is also almost as fast as the signing process.
// It’s also fast to perform batch signature verification with Ed25519.
// It’s built to be collision resilence. Hash-function collision won’t break the system.
func createED25519Key() error {
	const (
		kf = "SSH PK have been found"                                 // Key Found
		nk = "No SSH PK have been found, processing autoGenerateKey:" // No Key
		kg = "The SSH PK have been generated"                         // Key Generated
		ps = string(os.PathSeparator)                                 // Path Separator
	)
	// Returns expected ssh key: $HOME/.ssh/id_ed25519 or error
	sshKeyFile, ed25519 := sshKeyFile()
	if ed25519 != nil {
		log.Println(ed25519)
		return ed25519
	}
	// Check if file exist
	if _, err := os.Stat(sshKeyFile); os.IsNotExist(err) {
		// does not exist
		log.Printf("%s \n", nk)
		autoGenerateKey := Executable{
			Name:   "ssh-keygen",
			Params: []string{"-t", "ed25519", "-N", "", "-C", "tlsproxy", "-f", sshKeyFile},
		}
		// autogenerate the key
		autoGenerateKey.Run()
		log.Printf("%s \n", kg)
	} else {
		// exist
		log.Printf("%s \n", kf)
	}

	return nil
}
