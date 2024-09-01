package messages

import (
	"fmt"
	"time"
)

// Message holds the struct for new message on a particular channel
type Message struct {
	Channel   string
	Message   string
	Next      *Message
	Timestamp time.Time
	User      *Users
}

func (m *Message) String() string {
	return fmt.Sprintf("%s %s: %s", m.Timestamp.Format(time.RFC3339), m.User.Name, m.Message)
}

func NewMessage(user *Users, message, channel string, timestamp time.Time) *Message {
	return &Message{User: user, Message: message, Channel: channel, Timestamp: timestamp}
}

// AddMessage adds a new message to the linked list
func (m *Message) AddMessage(message *Message) {
	m.Next = message
}

// RemoveMessage  a -> b -> c
func (m *Message) RemoveMessage() {
	var next *Message = nil
	if m.Next.Next != nil {
		next = m.Next.Next
	} else {
		m.Next = next
	}
}

func (m *Message) MarshalBinary() ([]byte, error) {
	return nil, nil
}
