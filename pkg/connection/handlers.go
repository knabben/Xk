package connection

import (
	"fmt"
	"strings"
)

// Login set the nick and user information into the server.
func (i *IRCClient) Login() error {
	// Set the NICK command
	if err := i.Send(fmt.Sprintf("NICK %s\r\n", i.Nick)); err != nil {
		return err
	}
	// Set the USER information
	return i.Send(fmt.Sprintf("USER %s 0 * :%s\r\n", i.Nick, i.Name))
}

// Pong a handler for the PING command.
func (i *IRCClient) Pong(ping string) error {
	server := strings.Split(ping, ":")[1]
	return i.Send(fmt.Sprintf("PONG %s\r\n", server))
}
