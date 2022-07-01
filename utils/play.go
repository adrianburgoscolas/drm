package utils

import (
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
	"os"
	"time"
)

func Play(s string) error {
	f, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, ready, err := oto.NewContext(d.SampleRate(), 2, 2)
	if err != nil {
		return err
	}
	<-ready

	p := c.NewPlayer(d)
	defer p.Close()
	p.Play()

	for {
		time.Sleep(time.Second)
		if !p.IsPlaying() {
			break
		}
	}

	return nil
}
