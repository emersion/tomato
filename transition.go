package tomato

type TransitionFunc func(word string) int

// A transition to a state.
type Transition interface {
  Recognize(word string) (next *State, size int)
}

type transition struct {
  f TransitionFunc
  target *State
}

func (t *transition) Recognize(word string) (next *State, size int) {
  size = t.f(word)
  if size >= 0 {
    next = t.target
  }
  return
}

func newTransition(f TransitionFunc, target *State) Transition {
  return &transition{f, target}
}
