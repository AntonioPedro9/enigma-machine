package entities

type Plugboard struct {
	Mapping [26]rune
}

func NewPlugboard(mapping string) *Plugboard {
	var mappingRunes [26]rune
	copy(mappingRunes[:], []rune(mapping))

	return &Plugboard{
		Mapping: mappingRunes,
	}
}

func (p *Plugboard) Swap(char rune) rune {
	index := char - 'A'
	returnedChar := p.Mapping[index]
	return returnedChar
}
