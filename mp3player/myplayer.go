package main

import (
	"bufio"
	"fmt"
	"library"
	"mp"
	"os"
	"strconv"
	"strings"
	"time"
)

var lib *mp3player.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		{
			if len(tokens) == 6 {
				id++
				lib.Add(&mp3player.MusicEntry{strconv.Itoa(id),
					tokens[2], tokens[3], tokens[4], tokens[5]})
			} else {
				fmt.Println("USAGE: lib add <name><artist><source><type>")
			}
		}
	case "remove":
		if len(tokens) == 3 {
			_, index := lib.Find(tokens[2])
			lib.Remove(index)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}
	default:
		fmt.Println("Unrecognized lib command:", tokens[1])
	}
}
func handleCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println(`USAGE: play <name>
			stop <name>
			`)
		return
	}
	e, _ := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music", tokens[1], "does not exist.")
		return
	}
	switch tokens[0] {
	case "play":
		mp.Play(e.Source, e.Type, "play")
	case "stop":
		mp.Play(e.Source, e.Type, "stop")
	default:
		fmt.Println(`USAGE: play <name>
			stop <name>
			`)
	}
}
func main() {
	fmt.Println(`
￼￼￼Enter following commands to control the player:
lib list -- View the existing music lib
lib add <name><artist><source><type> -- Add a music to the music lib
lib remove <name> -- Remove the specified music from the lib
play <name> -- Play the specified music
`)
	lib = mp3player.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command-> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" || tokens[0] == "stop" {
			handleCommand(tokens)
			time.Sleep(time.Millisecond * 100)
		} else {
			fmt.Println("Unrecognized command:", tokens[0])
		}
	}
}
