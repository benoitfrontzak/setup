// The app purpose is to set up all required configuration to run tlsproxy as a service in 'another app'
// tlsproxy (ssh -R 80:localhost:8888 localhost.run)
// Requirements:
// SSH needs public key to run the tunnel, we're going to use ed25519 algorithm
// When run as a service tlsproxy must returns a file containing its configuration (to be read by 'another app')

package main

import "log"

func main() {
	setLog()

	// $HOME/.ssh/id_ed25519
	ed25519 := createED25519Key()
	if ed25519 != nil {
		log.Println(ed25519)
	}

	// $HOME/.ssh/.tunneled
	tunneled := createTunneledFile()
	if tunneled != nil {
		log.Println(tunneled)
	}

	// export bin
	addPath := exportBin()
	if addPath != nil {
		log.Println(addPath)
	}
}
