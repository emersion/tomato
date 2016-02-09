package test

import (
  "testing"

  "github.com/emersion/tomato"
  "github.com/emersion/tomato/transition"
)

// A simple automaton that recognizes words containing "ab".
func TestContains(t *testing.T) {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q2 := tomato.NewState()

  q0.Add(transition.NewRune('a', q1))
  q0.Add(transition.NewRune('b', q0))
  q0.Add(transition.NewRune('c', q0))

  q1.Add(transition.NewRune('a', q1))
  q1.Add(transition.NewRune('b', q2))
  q1.Add(transition.NewRune('c', q0))

  q2.Add(transition.NewRune('a', q2))
  q2.Add(transition.NewRune('b', q2))
  q2.Add(transition.NewRune('c', q2))

  a := tomato.NewAutomaton(q0, []*tomato.State{q2})

  items := []testCase{
    {"ab", true},
    {"abc", true},
    {"cab", true},
    {"cabc", true},
    {"aaaab", true},
    {"aba", true},
    {"abab", true},
    {"acab", true},
    {"a", false},
    {"b", false},
    {"aa", false},
    {"ba", false},
    {"bbbaaa", false},
    {"aca", false},
  }

  testAll(t, a, items)
}
