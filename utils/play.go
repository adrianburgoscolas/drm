package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

func InitPad(padFile string) (oto.Player, error) {
	f, err := os.Open(padFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return nil, err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return nil, err
	}
	<-ready
	pad := c.NewPlayer(d)
	return pad, nil
}

func Play(pad oto.Player) {
	pad.Play()
	fmt.Println("playing")
	for {
		time.Sleep(time.Second)
		if !pad.IsPlaying() {
			break
		}
	}
}
