package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

type WAVPlayer struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) chan string {

	ch := make(chan string, 0)
	go func(c chan string) {
		fmt.Println("Playing MP3 music", source)
		p.progress = 0
		for p.progress < 100 {
			if request := <-c; request == "stop" {
				fmt.Println("Music is going to be stopped...")
				break
			}
			time.Sleep(100 * time.Millisecond) // 假装正在播放 fmt.Print(".")
			// fmt.Println(".")
			p.progress += 1
		}
		fmt.Println("\nFinished playing", source)
	}(ch)
	return ch
}

func (p *WAVPlayer) Play(source string) chan string {
	ch := make(chan string, 0)
	fmt.Println("Playing WAV music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 假装正在播放 fmt.Print(".")
		p.progress += 1
	}
	fmt.Println("\nFinished playing", source)
	return ch
}
