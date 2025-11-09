package fastcaddy

import "golang.org/x/crypto/ssh"

// Option is a functional option for configuring the FastCaddy client.
type Option func(*FastCaddy)

// WithSSHClient provides an SSH client to tunnel Caddy API requests through.
func WithSSHClient(client *ssh.Client) Option {
	return func(fc *FastCaddy) {
		fc.sshClient = client
	}
}
