package operation

import (
  "github.com/emersion/tomato"
)

type Closure []*tomato.State

func (c Closure) contains(state *tomato.State) bool {
  for _, s := range c {
    if state == s {
      return true
    }
  }
  return false
}

func (c Closure) equals(other Closure) bool {
  if len(c) != len(other) {
    return false
  }

  for _, state := range c {
    if !other.contains(state) {
      return false
    }
  }

  return true
}

func epsilonClosure(state *tomato.State, closure Closure) Closure {
  closure = append(closure, state)

  for _, tr := range state.Transitions() {
    next, size := tr.Recognize("")
    if next == nil || size < 0 {
      continue
    }

    if closure.contains(next) {
      continue
    }

    nextClosure := epsilonClosure(next, closure)
    closure = append(closure, nextClosure...)
  }

  return closure
}

func EpsilonClosure(state *tomato.State) (closure Closure) {
  return epsilonClosure(state, Closure{})
}
