package entities

import "strings"

type Enigma struct {
	Rotors    []*Rotor
	Reflector *Reflector
	Plugboard *Plugboard
}

func NewEnigma(rotors []*Rotor, reflector *Reflector, plugboard *Plugboard) *Enigma {
	return &Enigma{
		Rotors:    rotors,
		Reflector: reflector,
		Plugboard: plugboard,
	}
}

func (e *Enigma) EncodeMessage(message string) string {
	message = strings.ToUpper(message)
	var encodedMessage strings.Builder

	rotor1 := e.Rotors[0]
	rotor2 := e.Rotors[1]
	rotor3 := e.Rotors[2]
	reflector := e.Reflector
	plugboard := e.Plugboard

	for _, char := range message {
		if char < 'A' || char > 'Z' {
			encodedMessage.WriteRune(char)
			continue
		}

		if rotor2.Position == rotor2.Notch {
			rotor2.Rotate()
			rotor3.Rotate()
		}
		if rotor1.Position == rotor1.Notch {
			rotor2.Rotate()
		}
		rotor1.Rotate()
	
		char = plugboard.Swap(char)
		char = rotor1.EncodeChar(char, true)
		char = rotor2.EncodeChar(char, true)
		char = rotor3.EncodeChar(char, true)
		char = reflector.Reflect(char)
		char = rotor3.EncodeChar(char, false)
		char = rotor2.EncodeChar(char, false)
		char = rotor1.EncodeChar(char, false)
		char = plugboard.Swap(char)
	
		encodedMessage.WriteRune(char)
	}	

	return encodedMessage.String()
}
