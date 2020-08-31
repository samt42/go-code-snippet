package main

import ()

// Video
type Video interface {
	Play()
}

type Movie struct{}

func (*Movie) Play() {}

// MusicVideo
type MusicVideo interface {
	Play()
	PlaySound()
}

type Drama struct{}

func (*Drama) Play() {}
func (*Drama) PlaySound() {}

func main() {
	var v Video
	v = new(Movie)
	v.Play()

	var m Drama
	m = new(Drama)
	m.Play()
	m.PlaySound()
}
