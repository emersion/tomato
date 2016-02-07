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

  q0.AddFunc(transition.Rune('a'), q1)
  q0.AddFunc(transition.Rune('b'), q0)
  q0.AddFunc(transition.Rune('c'), q0)

  q1.AddFunc(transition.Rune('a'), q1)
  q1.AddFunc(transition.Rune('b'), q2)
  q1.AddFunc(transition.Rune('c'), q0)

  q2.AddFunc(transition.Rune('a'), q2)
  q2.AddFunc(transition.Rune('b'), q2)
  q2.AddFunc(transition.Rune('c'), q2)

  a := tomato.NewAutomaton(q0, []*tomato.State{q2})

  items := []struct{
    input string
    recognized bool
  }{
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

  for _, item := range items {
    ok, _ := a.Recognize(item.input)
    if item.recognized && !ok {
      t.Error("Word '"+item.input+"' not recognized")
    }
    if !item.recognized && ok {
      t.Error("Word '"+item.input+"' recognized")
    }
  }
}
