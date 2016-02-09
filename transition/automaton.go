package transition

import (
  "github.com/emersion/tomato"
)

type automatonTransition struct {
  automaton *tomato.Automaton
  targets []*tomato.State
}

func (tr *automatonTransition) Recognize(word string) (*tomato.State, int) {
  ok, size, path := tr.automaton.Recognize(word)
  if ok {
    end := path[len(path)-1]
    for i, s := range tr.automaton.Ends() {
      if end == s {
        return tr.targets[i], size
      }
    }
  }

  return nil, -1
}

func (tr *automatonTransition) Targets() []*tomato.State {
  return tr.targets
}

func (tr *automatonTransition) Duplicate(targets ...*tomato.State) tomato.Transition {
  if len(targets) != len(tr.automaton.Ends()) {
    panic("Automaton transition targets must have same length as automaton ending states")
  }
  return NewAutomaton(tr.automaton, targets)
}

func NewAutomaton(a *tomato.Automaton, targets []*tomato.State) tomato.Transition {
  if len(a.Ends()) != len(targets) {
    panic("Automaton transition targets must have same length as automaton ending states")
  }

  return &automatonTransition{a, targets}
}
