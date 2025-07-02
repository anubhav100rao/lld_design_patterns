package mediator

import "fmt"

// Mediator interface
type GameMediator interface {
	Join(p *Player)
	Broadcast(sender *Player, action string)
}

// Colleague
type Player struct {
	name   string
	server GameMediator
}

func NewPlayer(name string) *Player { return &Player{name: name} }
func (p *Player) Join(server GameMediator) {
	p.server = server
	server.Join(p)
}
func (p *Player) SendAction(action string) {
	fmt.Printf("[%s] %s\n", p.name, action)
	p.server.Broadcast(p, action)
}
func (p *Player) Receive(senderName, action string) {
	fmt.Printf("[%s ← %s] %s\n", p.name, senderName, action)
}

// ConcreteMediator
type GameServer struct {
	players map[string]*Player
}

func NewGameServer() *GameServer {
	return &GameServer{players: make(map[string]*Player)}
}

func (gs *GameServer) Join(p *Player) {
	gs.players[p.name] = p
}

func (gs *GameServer) Broadcast(sender *Player, action string) {
	for _, p := range gs.players {
		if p.name != sender.name {
			p.Receive(sender.name, action)
		}
	}
}

// Client
func RunMultiPlayerGameMediator() {
	server := NewGameServer()
	alice := NewPlayer("Alice")
	bob := NewPlayer("Bob")
	carol := NewPlayer("Carol")

	alice.Join(server)
	bob.Join(server)
	carol.Join(server)

	alice.SendAction("moves to (10,20)")
	bob.SendAction("attacks Goblin")
}
