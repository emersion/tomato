package tomato

// An automaton recognizes words.
type Automaton struct {
  root *State
  ends []*State
}

// Get this automaton's ending states.
func (a *Automaton) Ends() []*State {
  return a.ends
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

func (a *Automaton) recognize(current *State, word string) (bool, int, []*State) {
  if len(word) == 0 && a.isEnd(current) {
    return true, 0, []*State{current}
  }

  var hasNext bool
  var bestSize int
  var bestPath []*State

  for _, child := range current.transitions {
    next, size := child.Recognize(word)
    if next == nil || size < 0 {
      continue
    }

    ok, processed, rpath := a.recognize(next, word[size:])
    total := size + processed
    if ok && (!hasNext || total > bestSize) {
      hasNext = true
      bestSize = total
      bestPath = append(rpath, current)
    }
  }

  if hasNext {
    return true, bestSize, bestPath
  }

  return a.isEnd(current), 0, []*State{current}
}

// Recognize a word. A not recognized part can be remaining at the end.
// Return values:
// - True if a part of the word has been recognized
// - The size of the recognized part
// - The recognition path
func (a *Automaton) Recognize(word string) (bool, int, []*State) {
  ok, size, rpath := a.recognize(a.root, word)

  // Reverse path
  n := len(rpath)
  path := make([]*State, n)
  for i := 0; i < n; i++ {
    path[i] = rpath[n-i-1]
  }

  return ok, size, path
}

// Recognize a whole word.
func (a *Automaton) RecognizeWhole(word string) (bool, []*State) {
  ok, size, path := a.Recognize(word)
  ok = (ok && size == len(word))
  return ok, path
}

// Create a new automaton, with a starting state and some ending states.
func NewAutomaton(root *State, ends []*State) *Automaton {
  return &Automaton{
    root: root,
    ends: ends,
  }
}
