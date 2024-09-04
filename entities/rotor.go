package entities

type Rotor struct {
	Mapping  [26]rune
	Position rune
	Notch    rune
	Ring     rune
}

func NewRotor(mapping string, position, notch, ring rune) *Rotor {
	var mappingRunes [26]rune
	copy(mappingRunes[:], []rune(mapping))

	return &Rotor{
		Mapping:  mappingRunes,
		Position: position,
		Notch:    notch,
		Ring:     ring,
	}
}

func (r *Rotor) Rotate() {
	positionOffset := r.Position - 'A'
	newPosition := (positionOffset + 1) % 26
	r.Position = rune(newPosition) + 'A'
}

func (r *Rotor) EncodeChar(char rune, forward bool) rune {
	if forward {
		return encodeForward(r, char)
	} else {
		return encodeBackward(r, char)
	}
}

func encodeForward(r *Rotor, char rune) rune {
	offset := int(r.Position - r.Ring)
	charIndex := int(char - 'A')
	adjustedIndex := (charIndex + offset + 26) % 26
	encodedChar := r.Mapping[adjustedIndex]
	encodedIndex := int(encodedChar - 'A')
	newPositionIndex := (encodedIndex - offset + 26) % 26

	return rune(newPositionIndex) + 'A'
}

func encodeBackward(r *Rotor, char rune) rune {
	offset := int(r.Position - r.Ring)
	charIndex := int(char - 'A')
	adjustedChar := rune((charIndex+offset+26)%26 + 'A')

	for i, mappingChar := range r.Mapping {
		if mappingChar == adjustedChar {
			newPositionIndex := (i - offset + 26) % 26
			return rune(newPositionIndex) + 'A'
		}
	}

	return char
}
