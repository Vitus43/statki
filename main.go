package main

import (
	"fmt"

	"github.com/Vitus43/statki/pkg/strzal"
)

//   1 2
//  +---+
// A| | |
// B| | | y
//  +---+
//    x
// mapa = {
// klucz -> "A": [p, p, p,] <- wartosc,
// "B": [p, p, p,],
//}

func main() {
	fmt.Println("Statki")
	// p := models.NowaDomyslnaPlansza()
	// p.Wypisz()
	for {
		strzal.PobierzStrzal()
	}
}

//    1 2
//   +-+-+
// A | | |
//	 +-+-+
// B | | | y
//   +-+-+
//    x

// nie moze byc dlugosc inna niz 2 ok
// pierwszy znak musi byc litera z dozwolonych wiersz
// drugi znak musi byc liczba
