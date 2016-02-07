package test

import (
  "testing"

  "github.com/emersion/tomato"
)

type testCase struct {
  input string
  recognized bool
}

func testAll(t *testing.T, automaton *tomato.Automaton, items []testCase) {
  for _, item := range items {
    ok, _ := automaton.RecognizeWhole(item.input)

    if item.recognized && !ok {
      t.Error("Word '"+item.input+"' not recognized")
    }
    if !item.recognized && ok {
      t.Error("Word '"+item.input+"' recognized")
    }
  }
}
