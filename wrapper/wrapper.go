package wrapper

import (
  "github.com/emersion/tomato"
  "github.com/emersion/tomato/transition"
)

func New(endsNbr int) (*tomato.Automaton, func(*tomato.Automaton)) {
  start := tomato.NewState()

  ends := make([]*tomato.State, endsNbr)
  for i := 0; i < endsNbr; i++ {
    ends[i] = tomato.NewState()
  }

  return tomato.NewAutomaton(start, ends), func (automaton *tomato.Automaton) {
    start.Reset()
    start.Add(transition.NewAutomaton(automaton, ends))
  }
}
