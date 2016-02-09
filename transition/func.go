package transition

import (
  "github.com/emersion/tomato"
)

type Func func(word string) int

type funcTransition struct {
  f Func
  target *tomato.State
}

func (tr *funcTransition) Recognize(word string) (next *tomato.State, size int) {
  size = tr.f(word)
  if size >= 0 {
    next = tr.target
  }
  return
}

func (tr *funcTransition) Targets() []*tomato.State {
  return []*tomato.State{tr.target}
}

func (tr *funcTransition) Duplicate(targets ...*tomato.State) tomato.Transition {
  if len(targets) != 1 {
    panic("A function transition needs exactly one target")
  }
  return NewFunc(tr.f, targets[0])
}

func NewFunc(f Func, target *tomato.State) tomato.Transition {
  return &funcTransition{f, target}
}
