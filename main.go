package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/knabben/Xk/pkg/connection"
	"github.com/knabben/Xk/pkg/messages"
	"log"
	"os"
	"time"
)

func main() {
	var err error
	nick := flag.String("nick", "nick", "Nickname")
	name := flag.String("name", "name", "Name")
	server := flag.String("server", "localhost:6667", "IRC server")
	flag.Parse()

	conn := connection.NewIRCClient(*nick, *name, *server)
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
	for {
		// iterate on stdin and send text via socket
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Scanln(text)
		if len(text) > 1 {
			cmd := fmt.Sprintf("PRIVMSG #alert :%s", text)
			if err := conn.Send(cmd); err != nil {
				log.Fatal(err)
			}
			user := messages.NewUsers(*nick)
			message := messages.NewMessage(user, text, "#alert", time.Now())
			conn.SaveMessage(message)
			fmt.Println(fmt.Sprintf("%s", message.String()))
		}
	}
}
