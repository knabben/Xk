package main

import (
	"flag"
	"github.com/knabben/Xk/pkg/connection"
	"log"
)

func main() {
	var err error
	nick := flag.String("nick", "nick", "Nickname")
	name := flag.String("name", "name", "Name")
	server := flag.String("server", "localhost:6667", "IRC server")
	flag.Parse()

	conn := connection.NewIRCClient(*nick, *name,
		*server)
	if err = conn.Connect(); err != nil {
		log.Fatal(err)
	}
	// Goroutine to parse all received messages.
	// should go to UI component or stdout
	go func() {
		err := conn.Receive()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if err = conn.Login(); err != nil {
		log.Fatal(err)
	}

	select {}
}
