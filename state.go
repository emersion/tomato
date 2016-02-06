package tomato

// An automaton's state.
// Contains transitions to reach children states.
type State struct {
  transitions []Transition
}

// Add a transition from this state to another one.
func (s *State) Add(tr Transition) {
  s.transitions = append(s.transitions, tr)
}

// Create a new state.
func NewState() *State {
  return &State{}
}
