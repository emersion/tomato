package transition

import (
  "github.com/emersion/tomato"
)

// Create a transition function that recognizes the empty word.
func Epsilon() Func {
  return func (word string) int {
    return 0
  }
}

// Create a transition that recognizes the empty word.
func NewEpsilon(target *tomato.State) tomato.Transition {
  return NewFunc(Epsilon(), target)
}
