package command

import "fmt"

// Command interface
type Command interface {
	Execute()
	Undo()
}

// Receiver
type Light struct {
	isOn bool
}

func (l *Light) On() {
	l.isOn = true
	fmt.Println("Light is ON")
}
func (l *Light) Off() {
	l.isOn = false
	fmt.Println("Light is OFF")
}

// Concrete Commands
type LightOnCommand struct{ light *Light }

func (c *LightOnCommand) Execute() { c.light.On() }
func (c *LightOnCommand) Undo()    { c.light.Off() }

type LightOffCommand struct{ light *Light }

func (c *LightOffCommand) Execute() { c.light.Off() }
func (c *LightOffCommand) Undo()    { c.light.On() }

// Invoker
type RemoteControl struct {
	onCommand  Command
	offCommand Command
	history    []Command
}

func (r *RemoteControl) SetCommands(on, off Command) {
	r.onCommand, r.offCommand = on, off
}

func (r *RemoteControl) PressOn() {
	r.onCommand.Execute()
	r.history = append(r.history, r.onCommand)
}

func (r *RemoteControl) PressOff() {
	r.offCommand.Execute()
	r.history = append(r.history, r.offCommand)
}

func (r *RemoteControl) PressUndo() {
	if len(r.history) == 0 {
		fmt.Println("Nothing to undo")
		return
	}
	cmd := r.history[len(r.history)-1]
	r.history = r.history[:len(r.history)-1]
	cmd.Undo()
}

// Client
func RunRemoteControl() {
	light := &Light{}
	onCmd := &LightOnCommand{light}
	offCmd := &LightOffCommand{light}

	remote := &RemoteControl{}
	remote.SetCommands(onCmd, offCmd)

	remote.PressOn()   // Light is ON
	remote.PressOff()  // Light is OFF
	remote.PressUndo() // Undo off → Light is ON
	remote.PressUndo() // Undo on  → Light is OFF
}
