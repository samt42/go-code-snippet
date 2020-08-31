package main

import "fmt"

type MovieFunc func(m *Movie)

type Movie struct{
	Name string
}

func (m *Movie) Play() {
	fmt.Println("Play Movie")
}

func (m *Movie) Execute(handler MovieFunc) {
	handler(m)
}

type MoviePlus struct {
	*Movie
	Description string
}

func MoviePlusFunc(m *MoviePlus) {
    m.Play()
}

func main() {
	m := new(Movie)
	m.Execute(MoviePlusFunc)
}
