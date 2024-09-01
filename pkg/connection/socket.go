package connection

import (
	"fmt"
	"github.com/knabben/Xk/pkg/messages"
	"log"
	"net"
	"regexp"
	"strings"
	"time"
)

const privRegex = `^:(?P<Nickname>[^\s!]+)!(?P<Username>[^\s@]+)@(?P<Host>[^\s]+) PRIVMSG (?P<Channel>[^\s]+) :(?P<Message>.+)$`

type Client interface {
	Connect() error
	Login() error
	Send(cmd string) error
	Join(channel string) error
	Receive() error
	SaveMessage(message *messages.Message)
}

type IRCClient struct {
	Nick        string
	Name        string
	Server      string
	LastMessage *messages.Message
	Connection  *net.Conn
}

func NewIRCClient(nick, name, server string) Client {
	return &IRCClient{
		Nick:       nick,
		Name:       name,
		Server:     server,
		Connection: nil,
	}
}

func (i *IRCClient) Connect() error {
	conn, err := net.Dial("tcp", i.Server)
	if err != nil {
		return err
	}
	i.Connection = &conn
	return nil
}

func (i *IRCClient) Send(cmd string) error {
	var conn = *i.Connection
	if _, err := conn.Write([]byte(cmd + "\r\n")); err != nil {
		return err
	}
	return nil
}

func (i *IRCClient) Join(channel string) error {
	return i.Send("JOIN #" + channel)
}

func (i *IRCClient) Receive() error {
	for {
		buf := make([]byte, 4096)
		n, err := (*i.Connection).Read(buf)
		if err != nil {
			return err
		}
		out := strings.Split(string(buf[:n]), "\r\n")
		for _, r := range out {
			switch {
			case strings.HasPrefix(r, "PING"):
				if err := i.Pong(r); err != nil {
					return err
				}
			case strings.Contains(r, "End of message of the day."):
				if err = i.Join("alert"); err != nil {
					log.Fatal(err)
				}
			case strings.Contains(r, "PRIVMSG"):
				param := params(r)
				message := messages.NewMessage(messages.NewUsers(param["Nickname"]), param["Message"], param["Channel"], time.Now())
				i.SaveMessage(message)
				fmt.Println(fmt.Sprintf("%s", message.String()))

			case strings.Contains(r, "NOTICE * "):
				fmt.Println(r)
			default:
			}
		}
	}
}

func (i *IRCClient) SaveMessage(message *messages.Message) {
	if i.LastMessage != nil {
		i.LastMessage.AddMessage(message)
	}
	i.LastMessage = message
}

func params(r string) (mapper map[string]string) {
	re := regexp.MustCompile(privRegex)
	match := re.FindStringSubmatch(r)
	mapper = make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			mapper[name] = match[i]
		}
	}
	return
}
