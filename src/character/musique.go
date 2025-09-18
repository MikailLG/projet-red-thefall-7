package character

import (
	"log"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

var ctrl *beep.Ctrl

func PlayMusic(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("Impossible de trouver le fichier son :", path)
		return
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Println("Erreur lors du décodage du MP3 :", err)
		return
	}
	loop := beep.Loop(-1, streamer)
	ctrl = &beep.Ctrl{Streamer: loop, Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0.0,
		Silent:   false,
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	go speaker.Play(volume)
}

func StopMusic() {
	if ctrl != nil {
		speaker.Lock()
		ctrl.Paused = true
		speaker.Unlock()
	}
}
