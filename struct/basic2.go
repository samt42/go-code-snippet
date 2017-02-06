package main

import (
	"fmt"
)

type Video struct {
	time int
}
type Music struct {
	time int
}

type MusicVideo struct {
	Video
	Music
}

func main() {
	v := &MusicVideo{}
	v.Video.time = 1
	v.Music.time = 2
	fmt.Println(v.Video.time)
	fmt.Println(v.Music.time)
}
