package transition

import (
  "github.com/emersion/tomato"
)

// Create a transition function that recognizes the empty word.
func Epsilon() tomato.TransitionFunc {
  return func (word string) int {
    return 0
  }
}
