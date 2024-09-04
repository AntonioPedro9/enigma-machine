package entities

type Reflector struct {
	Mapping [26]rune
}

func NewReflector(mapping string) *Reflector {
	var mappingRunes [26]rune
	copy(mappingRunes[:], []rune(mapping))

	return &Reflector{
		Mapping: mappingRunes,
	}
}

func (r *Reflector) Reflect(char rune) rune {
	index := char - 'A'
	returnedChar := r.Mapping[index]
	return returnedChar
}
