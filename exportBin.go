package main

import (
	"log"
	"os"
)

func exportBin() error {
	// Get current $PATH
	path := os.Getenv("PATH")

	// Create Path To Export: $PWD/bin
	pte, err := pathToExport()
	if err != nil {
		log.Println(err)
		return err
	}

	// Create new $PATH
	np := path + ";" + pte

	log.Println("BEFORE")
	log.Println(np)

	// Export new $PATH
	os.Setenv("PATH", np)

	log.Println("AFTER")
	log.Println(np)
	return nil
}
