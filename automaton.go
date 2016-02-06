package tomato

// An automaton recognizes words.
type Automaton struct {
  root *State
  ends []*State
}

// Check if a state is an end for this automaton.
func (a *Automaton) isEnd(state *State) bool {
  if state == nil {
    return false
  }

  for i := 0; i < len(a.ends); i++ {
    if state == a.ends[i] {
      return true
    }
  }

  return false
}

func (a *Automaton) recognize(state *State, word string) *State {
  if len(word) == 0 {
    return state
  }

  for _, child := range state.transitions {
    next, size := child.Recognize(word)
    if next == nil {
      continue
    }

    last := a.recognize(next, word[size:])
    if last != nil {
      return last
    }
  }

  return nil
}

// Recognize a word.
// Returns nil if this word is not recognized.
func (a *Automaton) Recognize(word string) *State {
  state := a.recognize(a.root, word)

  if a.isEnd(state) {
    return state
  }

  return nil
}

// Create a new automaton, with a starting state and some ending states.
func NewAutomaton(root *State, ends []*State) *Automaton {
  return &Automaton{
    root: root,
    ends: ends,
  }
}
