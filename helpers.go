package main

import (
	"log"
	"os"
)

const ps = string(os.PathSeparator)

// Returns string $HOME or error
func homeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("can't retrieve home directory: %v", err)
		return "", err
	}
	return home, nil
}

// Returns string $HOME/.ssh/.tunneled or error
func sshTunneledInfoFile() (string, error) {
	homeDirectory, err := homeDir()
	if err != nil {
		return "", err
	}
	sshTunneledInfo := homeDirectory + ps + ".ssh" + ps + ".tunneled"
	return sshTunneledInfo, nil
}

// Returns string $HOME/.ssh/id_ed25519 or error
func sshKeyFile() (string, error) {
	homeDirectory, err := homeDir()
	if err != nil {
		return "", err
	}
	sshKeyFile := homeDirectory + ps + ".ssh" + ps + "id_ed25519"
	return sshKeyFile, nil
}

// Returns string $PWD/bin or error
func pathToExport() (string, error) {
	// get $PWD
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return "", err
	}
	pathToExport := pwd + ps + "bin"
	return pathToExport, nil
}
