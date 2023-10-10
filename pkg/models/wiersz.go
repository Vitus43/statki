package models

import "fmt"

type Wiersz struct {
	Statki []*Statek
	Numer string
}

func (w Wiersz) Wypisz() {
	// TODO: przejsc w petli po statkach, sprawdzac jaka maja dlugosc i na jej podstawie wypisywac
	s := fmt.Sprintf("%s |", w.Numer)
	liczbaStatkow := len(w.Statki)
	for i := 0; i < liczbaStatkow; i++ {
		s += " |"
	}
	fmt.Println(s)
}