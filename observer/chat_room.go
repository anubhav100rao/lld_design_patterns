package observer

import "fmt"

// Observer interface
type Participant interface {
	Receive(sender, message string)
}

// Subject interface
type ChatRoom interface {
	Join(p Participant)
	Leave(p Participant)
	Broadcast(sender, msg string)
}

// ConcreteSubject
type SimpleChatRoom struct {
	participants map[Participant]struct{}
}

func NewChatRoom() *SimpleChatRoom {
	return &SimpleChatRoom{participants: make(map[Participant]struct{})}
}

func (cr *SimpleChatRoom) Join(p Participant) {
	cr.participants[p] = struct{}{}
}

func (cr *SimpleChatRoom) Leave(p Participant) {
	delete(cr.participants, p)
}

func (cr *SimpleChatRoom) Broadcast(sender, msg string) {
	for p := range cr.participants {
		if pName, ok := p.(*User); ok && pName.name == sender {
			continue // don't send to the sender
		}
		p.Receive(sender, msg)
	}
}

// ConcreteObserver
type User struct {
	name string
	room ChatRoom
}

func NewUser(name string) *User {
	return &User{name: name}
}

func (u *User) Receive(sender, msg string) {
	fmt.Printf("[%s] %s: %s\n", u.name, sender, msg)
}

func (u *User) Send(msg string) {
	u.room.Broadcast(u.name, msg)
}

// Client
func RunChatRoomObserver() {
	room := NewChatRoom()

	alice := NewUser("Alice")
	bob := NewUser("Bob")
	carol := NewUser("Carol")

	// assign room and join
	for _, u := range []*User{alice, bob, carol} {
		u.room = room
		room.Join(u)
	}

	alice.Send("Hello everyone!")
	bob.Send("Hey Alice!")
	room.Leave(carol)
	alice.Send("Where did Carol go?")
}
