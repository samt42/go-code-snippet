package main

import (
	"fmt"
)

type Language struct {
	Name string
}

func (l Language) Best() {
	fmt.Println(l.Name + "是最好的语言")
}

type AnotherLanguage struct {
	Language
}

func (l AnotherLanguage) Best() {
	fmt.Println(l.Name + "是最最好的语言")
}
func main() {
	v := AnotherLanguage{
		Language.Name: "PHP",
	}
	v.Name = "PHP"
	v.Best()
}
