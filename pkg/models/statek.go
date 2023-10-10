package models

type Statek struct {
	Dlugosc     int
	Wspolrzedne []*Wspolrzedna
	Zatopiony	bool
}

func (s *Statek) CzyZatopiony() {
	for _, wsp := range s.Wspolrzedne {
		if !wsp.Trafiona {
			return
		}
	}

	s.Zatopiony = true
}
