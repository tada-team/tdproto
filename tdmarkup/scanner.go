package tdmarkup

import (
	"fmt"
	"strings"
)

const EOF = rune(0)

func NewScanner(text string) *Scanner {
	return &Scanner{runes: []rune(text)}
}

type Scanner struct {
	runes []rune
	pos   int
}

// Current position
func (s Scanner) Position() int { return s.pos }

// Length in runes
func (s Scanner) Length() int { return len(s.runes) }

// Change current position
func (s *Scanner) Rewind(i int) {
	if i < 0 || i > s.Length() {
		panic(fmt.Errorf("index must be between 0 and %d, %d given", s.Length(), i))
	}
	s.pos = i
}

// Diff
func (s Scanner) Since(i int) int {
	if s.pos < i {
		panic(fmt.Errorf("index must be between less than %d, %d given", s.pos, i))
	}
	return s.pos - i
}

// Number of runes
func (s Scanner) Rest() int {
	return s.Length() - s.pos
}

// Previous rune
func (s Scanner) Prev() rune {
	if s.pos < 2 {
		return EOF
	}
	return s.runes[s.pos-2]
}

// Current rune: FIXME: strange logic, but it works
func (s Scanner) Current() rune {
	if s.pos < 1 {
		return EOF
	}
	return s.runes[s.pos-1]
}

// Returns next rune without rewind
func (s Scanner) Next() rune {
	if s.pos >= s.Length() {
		return EOF
	}
	return s.runes[s.pos]
}

// Returns next rune and rewind current position
func (s *Scanner) TakeNext() rune {
	ch := s.Next()
	if ch != EOF {
		s.pos++
	}
	return ch
}

func (s *MarkupScanner) ScanUntil(cl []rune) string {
	start := s.Position()
	var b strings.Builder

	switch len(cl) {
	case 0:
		panic("invalid close operator length: 0")
	case 1:
		r := cl[0]
		for s.Rest() > 0 {
			b.WriteRune(s.TakeNext())
			if s.Current() == r {
				return b.String()
			}
		}
		s.Rewind(start)
		return ""
	default:
		if s.Rest() < len(cl) {
			return ""
		}
		for _, r := range cl {
			if s.Next() != r {
				s.Rewind(start)
				return ""
			}
			b.WriteRune(s.TakeNext())
		}
	}

	return b.String()
}
