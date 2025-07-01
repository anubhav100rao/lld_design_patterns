package adapter

import "fmt"

// Target interface
type MediaPlayer interface {
	Play(audioType, fileName string)
}

// Adaptee interface
type AdvancedMediaPlayer interface {
	PlayVLC(fileName string)
	PlayMP4(fileName string)
}

// Concrete Adaptees
type VLCPlayer struct{}

func (v *VLCPlayer) PlayVLC(fileName string) {
	fmt.Println("Playing vlc file. Name:", fileName)
}
func (v *VLCPlayer) PlayMP4(fileName string) {}

type MP4Player struct{}

func (m *MP4Player) PlayMP4(fileName string) {
	fmt.Println("Playing mp4 file. Name:", fileName)
}
func (m *MP4Player) PlayVLC(fileName string) {}

// Adapter
type MediaAdapter struct {
	advancedPlayer AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	var player AdvancedMediaPlayer
	switch audioType {
	case "vlc":
		player = &VLCPlayer{}
	case "mp4":
		player = &MP4Player{}
	}
	return &MediaAdapter{advancedPlayer: player}
}

func (m *MediaAdapter) Play(audioType, fileName string) {
	switch audioType {
	case "vlc":
		m.advancedPlayer.PlayVLC(fileName)
	case "mp4":
		m.advancedPlayer.PlayMP4(fileName)
	}
}

// Client – original simple player
type AudioPlayer struct {
	adapter *MediaAdapter
}

func (a *AudioPlayer) Play(audioType, fileName string) {
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file. Name:", fileName)
	} else if audioType == "vlc" || audioType == "mp4" {
		a.adapter = NewMediaAdapter(audioType)
		a.adapter.Play(audioType, fileName)
	} else {
		fmt.Println("Invalid media. ", audioType, " format not supported")
	}
}

func RunMusicAdapterDemo() {
	player := &AudioPlayer{}

	player.Play("mp3", "song.mp3")
	player.Play("mp4", "video.mp4")
	player.Play("vlc", "movie.vlc")
	player.Play("avi", "mind.avi")
}
