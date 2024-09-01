package messages

import (
	"fmt"
	"log"
	"strings"
	"testing"
	"time"
)

func TestStructRender(t *testing.T) {
	user := NewUsers("fake")
	message := NewMessage(user, "new message", "#chan", time.Now())
	output := fmt.Sprintf("%s", message)
	expected := "fake: new message"
	if strings.Contains(expected, output) {
		log.Fatalf("error %s", message)
	}
}

func TestLLAddNewMessage(t *testing.T) {
	user := NewUsers("fake")
	message1 := NewMessage(user, "a new message 1", "#chan", time.Now())
	message2 := NewMessage(user, "a new message 2", "#chan", time.Now())
	message1.AddMessage(message2)
	if message1.Next != message2 {
		log.Fatalf("error trying to fetch new message")
	}
}

func TestLLRemove1Message(t *testing.T) {
	user := NewUsers("fake")
	message1 := NewMessage(user, "a new message 1", "#chan", time.Now())
	message2 := NewMessage(user, "a new message 2", "#chan", time.Now())
	message3 := NewMessage(user, "a new message 3", "#chan", time.Now())
	message1.AddMessage(message2)
	message2.AddMessage(message3)
	message1.RemoveMessage()
	if message1.Next == message3 {
		log.Fatalf("error removing the message")
	}
}

func TestLLRemoveLastMessage(t *testing.T) {
	user := NewUsers("fake")
	message1 := NewMessage(user, "a new message 1", "#chan", time.Now())
	message2 := NewMessage(user, "a new message 2", "#chan", time.Now())
	message1.AddMessage(message2)
	message1.RemoveMessage()
	if message1.Next == message2 {
		log.Fatalf("error removing the message")
	}
}
