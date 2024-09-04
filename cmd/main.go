package main

import (
	"fmt"
	"strings"

	"github.com/AntonioPedro9/enigma-machine/entities"
)

var (
	rotorI        = entities.NewRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'A', 'Q', 'A')
	rotorII       = entities.NewRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", 'A', 'E', 'A')
	rotorIII      = entities.NewRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", 'A', 'V', 'A')
	rotorIV       = entities.NewRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", 'A', 'J', 'A')
	rotorV        = entities.NewRotor("VZBRGITYUPSDNHLXAWMJQOFECK", 'A', 'Z', 'A')
	reflectorUKWB = entities.NewReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	reflectorUKWC = entities.NewReflector("AFBVCPDJEIGOHYKRLZMXNWTQSU")
	plugboard     = entities.NewPlugboard("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func main() {
	enigma := entities.NewEnigma([]*entities.Rotor{rotorI, rotorII, rotorIII}, reflectorUKWB, plugboard)

	message := "SECRET MESSAGE"
	message = strings.ToUpper(message)

	encodedMessage := enigma.EncodeMessage(message)

	fmt.Println(encodedMessage)
}
