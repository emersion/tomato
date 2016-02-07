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

// Add a transition from this state to another one through a specified function.
func (s *State) AddFunc(f TransitionFunc, target *State) {
  s.Add(newTransition(f, target))
}

// Create a new state.
func NewState() *State {
  return &State{}
}
