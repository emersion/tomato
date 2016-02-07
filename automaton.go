package tomato

import "log"

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

func (a *Automaton) recognize(state *State, word string) (bool, []*State) {
  if len(word) == 0 && a.isEnd(state) {
    return true, []*State{state}
  }

  for _, child := range state.transitions {
    next, size := child.Recognize(word)
    if next == nil {
      continue
    }
    log.Println(word, size)
    ok, path := a.recognize(next, word[size:])
    if ok {
      return true, append(path, state)
    }
  }

  return false, []*State{state}
}

// Recognize a word.
// Returns nil if this word is not recognized.
func (a *Automaton) Recognize(word string) (bool, []*State) {
  ok, rpath := a.recognize(a.root, word)

  // Reverse path
  n := len(rpath)
  path := make([]*State, n)
  for i := 0; i < n; i++ {
    path[i] = rpath[n-i-1]
  }

  return ok, path
}

// Create a new automaton, with a starting state and some ending states.
func NewAutomaton(root *State, ends []*State) *Automaton {
  return &Automaton{
    root: root,
    ends: ends,
  }
}
