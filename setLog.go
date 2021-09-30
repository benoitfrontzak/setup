package main

import "log"

func setLog() {
	log.SetPrefix("[INFO]: ")
	log.SetFlags(0) // remove file:line and timestamps from log liness
}
