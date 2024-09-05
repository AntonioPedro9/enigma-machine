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
	reflectorUKWB = entities.NewReflector("BR CU DH EQ FS GL IP JX KN MO TZ VW")
	reflectorUKWC = entities.NewReflector("AF BV CP DJ EI GO HY KR LZ MX NW TQ SU")
	plugboard     = entities.NewPlugboard("BQ CR DI EJ KW MT OS PX UZ GH")
)

func main() {
	enigma := entities.NewEnigma([]*entities.Rotor{rotorI, rotorII, rotorIII}, reflectorUKWB, plugboard)

	message := "secret message"
	message = strings.ToUpper(message)

	encodedMessage := enigma.EncodeMessage(message)

	fmt.Println(encodedMessage)
}
