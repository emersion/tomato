package transition

import (
  "github.com/emersion/tomato"
)

func String(value string) tomato.TransitionFunc {
  return func (word string) int {
    if len(word) < len(value) {
      return -1
    }

    if word[:len(value)] == value {
      return len(value)
    }

    return -1
  }
}
