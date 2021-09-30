# setup

Setup all required configuration to run tlsproxy as a service for the app 'chat-host'
It creates:
$HOME/.ssh/id_ed25519 (SSH needs public key to run the tunnel, we're going to use ed25519)
$HOME/.ssh/.tunneled (Configuration of tlsproxy : pid/termination)
export bin/tlsproxy to $PATH (TO DO...)
