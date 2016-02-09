package transition

import (
  "unicode/utf8"

  "github.com/emersion/tomato"
)

func Rune(value rune) Func {
  return func (word string) int {
    if len(word) == 0 {
      return -1
    }

    ch, size := utf8.DecodeRuneInString(word)
    if ch == value {
      return size
    }

    return -1
  }
}

func NewRune(value rune, target *tomato.State) tomato.Transition {
  return NewFunc(Rune(value), target)
}
