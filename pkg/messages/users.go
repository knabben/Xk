package messages

type Users struct {
	Name string
}

func NewUsers(name string) *Users {
	return &Users{name}
}
