package transition

import (
  "unicode/utf8"

  "github.com/emersion/tomato"
)

type runeTransition struct {
  value rune
  target *tomato.State
}

func (t *runeTransition) Recognize(word string) (*tomato.State, int) {
  if t.value == tomato.Epsilon {
    return t.target, 0
  }

  ch, size := utf8.DecodeRuneInString(word)
  if ch == t.value {
    return t.target, size
  }

  return nil, 0
}

// Create a new transition that recognizes a single character.
func NewRune(value rune, target *tomato.State) tomato.Transition {
  return &runeTransition{
    value: value,
    target: target,
  }
}
