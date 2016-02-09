package transition

import (
  "github.com/emersion/tomato"
)

func String(value string) Func {
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

func NewString(value string, target *tomato.State) tomato.Transition {
  return NewFunc(String(value), target)
}
