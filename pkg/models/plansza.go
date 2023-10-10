package models

import (
	"fmt"
	"sort"
)
type Plansza struct {
	Plansza map[string]Wiersz
}

const (
	a = "A"
	b = "B"
	c = "C"
)

func (p Plansza) Wypisz() {
	l := len(p.PobierzWiersz(a).Statki)
	s := "   "
	// wypisanie liczb kolumn
	for i := 1; i <= l; i++ {
		s += fmt.Sprintf("%d ", i)
	}
	fmt.Println(s)
	WypiszObramowke(l)
	wiersze := []Wiersz{}
	for _, wiersz := range p.Plansza{
		wiersze = append(wiersze, wiersz)
	}

	sort.Slice(wiersze, func(i, j int) bool {
		return wiersze[i].Numer < wiersze[j].Numer
	})

	for _, wiersz := range wiersze {
		wiersz.Wypisz()
		WypiszObramowke(l)
	}
}

func (p Plansza) DodajWiersz(nazwaWiersza string, wiersz Wiersz) {
	p.Plansza[nazwaWiersza] = wiersz
}

func (p Plansza) PobierzWiersz(nazwaWiersza string) Wiersz {
	wiersz, ok := p.Plansza[nazwaWiersza]
	if !ok {
		panic("wiersz nie istnieje")
	}

	return wiersz
}

func WypiszObramowke(l int) {
	s := "  +"
	for i := 0; i < l; i++ {
		s += "-+"
	}
	fmt.Println(s)
}

func NowaDomyslnaPlansza() Plansza {
	p := Plansza{map[string]Wiersz{}}
	p.DodajWiersz(a, Wiersz{
		Numer: a,
		Statki: []*Statek{
			{
				Dlugosc: 1,
				Wspolrzedne: []*Wspolrzedna{
					{
						X: 1,
						Y: a,
					},
				},
			},
			{
				Dlugosc: 1,
				Wspolrzedne: []*Wspolrzedna{
					{
						X: 2,
						Y: a,
					},
				},
			},
			{
				Dlugosc: 1,
				Wspolrzedne: []*Wspolrzedna{
					{
						X: 3,
						Y: a,
					},
				},
			},
		},
	})
	p.DodajWiersz(b, Wiersz{
		Numer: b,
		Statki: []*Statek{
			{
				Dlugosc: 3,
				Wspolrzedne: []*Wspolrzedna{
					{
						X: 1,
						Y: b,
					},
					{
						X: 2,
						Y: b,
					},
					{
						X: 3,
						Y: b,
					},
				},
			},
		},
	})

	return p
}