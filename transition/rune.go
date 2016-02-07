package transition

import (
  "unicode/utf8"

  "github.com/emersion/tomato"
)

func Rune(value rune) tomato.TransitionFunc {
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
