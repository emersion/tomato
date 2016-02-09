package operation

import (
  "github.com/emersion/tomato"
)

func closureTransitions(closure Closure) (transitions []tomato.Transition) {
  for _, state := range closure {
    for _, tr := range state.Transitions() {
      // TODO: remove targets already in closure
      inside := false
      for _, target := range tr.Targets() {
        if closure.contains(target) {
          inside = true
          break
        }
      }
      if inside {
        continue
      }

      transitions = append(transitions, tr)
    }
  }
  return
}

func closureIsEnd(closure Closure, ends []*tomato.State) bool {
  for _, end := range ends {
    if closure.contains(end) {
      return true
    }
  }
  return false
}

// Transform any automaton to a synchronized and determinized automaton.
func Determinize(a *tomato.Automaton) *tomato.Automaton {
  closures := []Closure{}
  states := []*tomato.State{}
  ends := []*tomato.State{}

  var collect func (state *tomato.State) *tomato.State
  collect = func (state *tomato.State) *tomato.State {
    closure := EpsilonClosure(state)

    // Make sure this closure hasn't been already processed
    for i, c := range closures {
      if closure.equals(c) {
        return states[i]
      }
    }

    // Create a new state for this closure
    state = tomato.NewState()

    // Collect this closure's transitons
    trs := closureTransitions(closure)

    // Add closure & new state to list
    closures = append(closures, closure)
    states = append(states, state)

    // Does this new state is an ending one?
    if closureIsEnd(closure, a.Ends()) {
      ends = append(ends, state)
    }

    // Add transitions to new state
    for _, tr := range trs {
      targets := tr.Targets()
      newTargets := make([]*tomato.State, len(targets))
      for i, target := range targets {
        newTargets[i] = collect(target)
      }
      state.Add(tr.Duplicate(newTargets...))
    }

    return state
  }

  root := collect(a.Root())

  return tomato.NewAutomaton(root, ends)
}
