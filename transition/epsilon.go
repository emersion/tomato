package transition

import (
  "github.com/emersion/tomato"
)

type epsilonTransition struct {
  target *tomato.State
}

func (t *epsilonTransition) Recognize(word string) (*tomato.State, int) {
  return t.target, 0
}

// Create a new transition that recognizes the empty word.
func NewEpsilon(target *tomato.State) tomato.Transition {
  return &epsilonTransition{
    target: target,
  }
}
