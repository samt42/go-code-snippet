package main

import ()

// Video
type Video interface {
	Play()
}

type Movie struct{}

func (*Movie) Play() {}

type Drama struct{}

func (*Drama) Play() {}

func main() {
	var v Video
	v = new(Movie)
	v.Play()

	v = new(Drama)
	v.Play()
}
