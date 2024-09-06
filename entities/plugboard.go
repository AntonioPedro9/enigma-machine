package entities

import "strings"

type Plugboard struct {
    Mapping [26]rune
}

func NewPlugboard(pairs string) *Plugboard {
	pairs = strings.ToUpper(pairs)
    var mappingRunes [26]rune

    for i := 0; i < 26; i++ {
        mappingRunes[i] = rune('A' + i)
    }

    for i := 0; i < len(pairs); i += 3 {
        if i+1 < len(pairs) && pairs[i+1] != ' ' {
            a := rune(pairs[i])
            b := rune(pairs[i+1])
            mappingRunes[a-'A'] = b
            mappingRunes[b-'A'] = a
        }
    }

    return &Plugboard{
        Mapping: mappingRunes,
    }
}

func (p *Plugboard) Swap(char rune) rune {
    index := char - 'A'
    returnedChar := p.Mapping[index]
    return returnedChar
}
