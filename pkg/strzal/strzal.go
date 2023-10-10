package strzal

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Vitus43/statki/pkg/models"
)

const (
	a = "A"
	b = "B"
	c = "C"
	d = "D"
	e = "E"
	f = "F"
	g = "G"
	h = "H"
)

func PobierzStrzal() *models.Wspolrzedna {
	var input string
	fmt.Println("Podaj strzal: ")
	fmt.Scan(&input)
	fmt.Println("Strzal:", input)
	fmt.Println(len(input))
	wsp, err := ZweryfikujStrzal(input)
	if err != nil {
		fmt.Printf("weryfikacja strzalu: %s\n", err)
		PobierzStrzal()
	}

	return wsp
}

// input od uzytkownika np. "a2"
func ZweryfikujStrzal(input string) (*models.Wspolrzedna, error) {
	if len(input) > 2 {
		return nil, fmt.Errorf("strzal musi skaldac sie z 2 znakow")
	}
	// podzielone lista stringow np. ["a", "2"]
	podzielone := strings.Split(input, "")
	if len(podzielone) != 2 {
		return nil, fmt.Errorf("strzal nie sklada sie z 2 znakow; strzal: %s", input)
	}

	wiersz, kolumna := strings.ToUpper(podzielone[0]), podzielone[1]
	err := ZweryfikujWiersz(wiersz)
	if err != nil {
		return nil, fmt.Errorf("weryfikacja wiersza: %w", err)
	}

	kol, err := ZweryfikujKolumne(kolumna)
	if err != nil {
		return nil, fmt.Errorf("weryfikacja kolumny: %w", err)
	}

	return &models.Wspolrzedna{X: kol, Y: wiersz}, nil
}

func ZweryfikujWiersz(wiersz string) error {
	literyWierszy := []string{a, b, c, d, e, f, g, h}

	dobry := false

	for _, lw := range literyWierszy {
		if lw == wiersz {
			dobry = true
			break
		}
	}

	if !dobry {
		return fmt.Errorf("pierwszy znak musi być literą wiersza, dostepne wiersze %s", literyWierszy)
	}

	return nil
}

func ZweryfikujKolumne(k string) (int, error) {
	kol, err := strconv.Atoi(k)
	if err != nil {
		return 0, fmt.Errorf("zamiana string na int: %w", err)
	}

	if kol == 0 {
		return 0, fmt.Errorf("kolumna musi byc miedzy 1-9")
	}

	return kol, nil
}
