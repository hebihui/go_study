package mp

import (
	"fmt"
)

type Player interface {
	Play(source string) chan string
}

//增加一个map用于保存，不同的chanel
var conns map[string]chan string = make(map[string]chan string)

func Play(source, mtype, cmd string) {
	var p Player
	if cmd == "play" {
		switch mtype {
		case "MP3":
			p = &MP3Player{}
		case "WAV":
			p = &WAVPlayer{}
		default:
			fmt.Println("unsupported music type", mtype)
			return
		}
		ch := p.Play(source)
		conns[source] = ch
	} else if cmd == "stop" {
		conns[source] <- "stop"
	}
}
