package facade

import "fmt"

// Subsystem components

type Amplifier struct{}

func (a *Amplifier) On() { fmt.Println("Amplifier on") }
func (a *Amplifier) SetVolume(v int) {
	fmt.Printf("Amplifier setting volume to %d\n", v)
}
func (a *Amplifier) Off() { fmt.Println("Amplifier off") }

type DVDPlayer struct{}

func (d *DVDPlayer) On() { fmt.Println("DVD Player on") }
func (d *DVDPlayer) Play(movie string) {
	fmt.Printf("DVD Player playing \"%s\"\n", movie)
}
func (d *DVDPlayer) Stop() { fmt.Println("DVD Player stopped") }
func (d *DVDPlayer) Off()  { fmt.Println("DVD Player off") }

type Projector struct{}

func (p *Projector) On() { fmt.Println("Projector on") }
func (p *Projector) WideScreenMode() {
	fmt.Println("Projector in widescreen mode")
}
func (p *Projector) Off() { fmt.Println("Projector off") }

// Facade

type HomeTheaterFacade struct {
	amp       *Amplifier
	dvd       *DVDPlayer
	projector *Projector
}

func NewHomeTheaterFacade(amp *Amplifier, dvd *DVDPlayer, proj *Projector) *HomeTheaterFacade {
	return &HomeTheaterFacade{amp: amp, dvd: dvd, projector: proj}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.amp.On()
	h.amp.SetVolume(5)
	h.projector.On()
	h.projector.WideScreenMode()
	h.dvd.On()
	h.dvd.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.dvd.Stop()
	h.dvd.Off()
	h.projector.Off()
	h.amp.Off()
}

// Client

func RunHomeTheaterDemo() {
	amp := &Amplifier{}
	dvd := &DVDPlayer{}
	projector := &Projector{}

	theater := NewHomeTheaterFacade(amp, dvd, projector)
	theater.WatchMovie("Inception")
	fmt.Println()
	theater.EndMovie()
}
