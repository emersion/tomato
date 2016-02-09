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

// Remove all transitions from this state.
func (s *State) Reset() {
  s.transitions = nil
}

// Get all transitions from this state.
func (s *State) Transitions() []Transition {
  return s.transitions
}

// Create a new state.
func NewState() *State {
  return &State{}
}
