package mediator

import "fmt"

// —— Mediator interface ——
type ChatMediator interface {
	SendMessage(sender, msg string)
	Register(user *User)
}

// —— ConcreteMediator ——
type ChatRoom struct {
	users map[string]*User
}

func NewChatRoom() *ChatRoom {
	return &ChatRoom{users: make(map[string]*User)}
}

func (cr *ChatRoom) Register(user *User) {
	cr.users[user.name] = user
	user.mediator = cr
}

func (cr *ChatRoom) SendMessage(sender, msg string) {
	for _, u := range cr.users {
		if u.name != sender {
			u.Receive(sender, msg)
		}
	}
}

// —— Colleague ——
type User struct {
	name     string
	mediator ChatMediator
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Send(msg string) {
	fmt.Printf("[%s ➜ room] %s\n", u.name, msg)
	u.mediator.SendMessage(u.name, msg)
}

func (u *User) Receive(sender, msg string) {
	fmt.Printf("[%s ⬅ %s] %s\n", u.name, sender, msg)
}

// —— Client ——
func RunChatRoomMediator() {
	room := NewChatRoom()

	alice := NewUser("Alice")
	bob := NewUser("Bob")
	carol := NewUser("Carol")

	room.Register(alice)
	room.Register(bob)
	room.Register(carol)

	alice.Send("Hi everyone!")
	bob.Send("Hey Alice!")
	carol.Send("Hello folks!")
}
