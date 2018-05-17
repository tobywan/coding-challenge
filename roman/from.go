// Package roman is a roman numeral converter
package roman

import (
	"errors"
	"fmt"
)

const (
	// M = 1000
	M = 1000
	// D = 500
	D = 500
	// C = 100
	C = 100
	// L = 50
	L = 50
	// X = 10
	X = 10
	// V = 5
	V = 5
	// I = 1
	I = 1
)

type numeral int

var (
	values = map[rune]numeral{
		'M': M,
		'D': D,
		'C': C,
		'L': L,
		'X': X,
		'V': V,
		'I': I,
	}
	subtractors = map[numeral]numeral{
		I: 0,
		V: I,
		X: I,
		L: X,
		C: X,
		D: C,
		M: C,
	}
)

const (
	maxOccurs = 3
)

// A sequence of units, 5s, 10s, 50s etc
type sequence struct {
	// the roman numeral numeric value
	num numeral
	// How many times repeated?
	occurs int
	// what the subtractor for this numeral is, e.g C is X
	// if there is not a subtractor, e.g. for I, then zero
	subtractor numeral
	// Whether we need to subtract the subtractor
	haveSubtract bool
}

type parser struct {
	sequences []sequence
}

func newSeq(n numeral) *sequence {
	return &sequence{
		num:        n,
		occurs:     1,
		subtractor: subtractors[n],
	}
}

func (p *parser) receive(n numeral) error {
	// Accept the numeral if:
	//
	l := len(p.sequences)
	if l == 0 {
		// Accept the first numeral
		p.sequences = append(p.sequences, *newSeq(n))
		return nil
	}
	s := &(p.sequences[l-1])
	// More of the same numeral. Allow maxOccurs + 1 if one has been used for subtraction
	if n == s.num {
		if s.occurs == maxOccurs && !s.haveSubtract {
			return errors.New("maximum occurences exceeded")
		}

		s.occurs++
		return nil
	}
	if n < s.num {
		if n != s.subtractor {
			return errors.New("invalid order")
		}
		if s.haveSubtract {
			return errors.New("multiple subtractors")
		}

		if s.occurs > 1 {
			return errors.New("subtracting from multiple occurrence")
		}
		// Has this subractor been used as an addor, e.g. IVI
		if l > 1 {
			prev := &(p.sequences[l-2])
			if n == prev.num && !prev.haveSubtract {
				return errors.New("reuse as subractor")
			}
			// Are we subtracting from a numeral used as an addor?

		}
		s.haveSubtract = true
		// Reset the occurrences as this now doesn't count
		//s.occurs = 0
		return nil
	}
	// It's a new larger number
	p.sequences = append(p.sequences, *newSeq(n))

	return nil
}

func (p *parser) calc() int {
	var result int
	for _, seq := range p.sequences {
		result += seq.occurs * int(seq.num)
		if seq.haveSubtract {
			result -= int(seq.subtractor)
		}
	}
	return result
}

// From converts a valid string comprising upper case
// M,D,C,L,X,V,I in valid order to an int. If not valid the error is non nil
func From(roman string) (int, error) {
	if len(roman) == 0 {
		return 0, nil
	}
	if err := validate(roman); err != nil {
		return 0, err
	}

	runes := []rune(roman)
	// Possiby over allocating length of sequences.
	p := parser{}

	for i := len(runes) - 1; i > -1; i-- {
		if err := p.receive(values[runes[i]]); err != nil {
			return 0, fmt.Errorf("unexpected numeral %q in %q: %v", runes[i], roman, err)
		}
	}

	return p.calc(), nil
}

func validate(roman string) error {
	for _, r := range roman {
		if _, ok := values[r]; !ok {
			return fmt.Errorf("invalid character %q in %q", r, roman)
		}
	}
	return nil
}
