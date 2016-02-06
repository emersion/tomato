package tomato

// A transition to a state.
type Transition interface {
  Recognize(word string) (next *State, size int)
}
