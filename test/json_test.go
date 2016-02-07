package test

import (
  "testing"

  "github.com/emersion/tomato"
  "github.com/emersion/tomato/transition"
)

func TestJsonNumber(t *testing.T) {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q1_1 := tomato.NewState()
  q2 := tomato.NewState()
  q3 := tomato.NewState()
  q4 := tomato.NewState()
  q5 := tomato.NewState()
  q5_1 := tomato.NewState()
  q5_2 := tomato.NewState()
  q5_3 := tomato.NewState()
  q6 := tomato.NewState()

  q0.Add(transition.NewRune('-', q1))
  q0.Add(transition.NewEpsilon(q1))

  q1.Add(transition.NewRune('0', q2))
  q1.Add(transition.NewRegexp("[1-9]", q1_1))

  q1_1.Add(transition.NewRegexp("[0-9]", q1_1))
  q1_1.Add(transition.NewEpsilon(q2))

  q2.Add(transition.NewEpsilon(q5))
  q2.Add(transition.NewRune('.', q3))

  q3.Add(transition.NewRegexp("[0-9]", q4))

  q4.Add(transition.NewRegexp("[0-9]", q4))
  q4.Add(transition.NewEpsilon(q5))

  q5.Add(transition.NewRegexp("[eE]", q5_1))
  q5.Add(transition.NewEpsilon(q6))

  q5_1.Add(transition.NewRegexp("[+-]?", q5_2))

  q5_2.Add(transition.NewRegexp("[0-9]", q5_3))

  q5_3.Add(transition.NewRegexp("[0-9]", q5_3))
  q5_3.Add(transition.NewEpsilon(q6))

  number := tomato.NewAutomaton(q0, []*tomato.State{q6})

  items := []struct{
    input string
    recognized bool
  }{
    {"0", true},
    {"1", true},
    {"01", false},
    {"1.0", true},
    {"786.632973", true},
    {"0.367283", true},
    {"0.0", true},
    {"67e68", true},
    {"16E98", true},
  }

  for _, item := range items {
    ok, _ := number.Recognize(item.input)
    if item.recognized && !ok {
      t.Error("Word '"+item.input+"' not recognized")
    }
    if !item.recognized && ok {
      t.Error("Word '"+item.input+"' recognized")
    }
  }
}

func TestString(t *testing.T) {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q2 := tomato.NewState()
  q2_1 := tomato.NewState()
  q2_1_1 := tomato.NewState()
  q3 := tomato.NewState()
  q4 := tomato.NewState()

  q0.Add(transition.NewRune('"', q1))

  q1.Add(transition.NewEpsilon(q2))
  q1.Add(transition.NewEpsilon(q3))

  q2.Add(transition.NewRegexp("[^\\\"]+", q3))
  q2.Add(transition.NewRune('\\', q2_1))

  q2_1.Add(transition.NewRegexp("[\"\\/bfnrt]", q3))
  q2_1.Add(transition.NewRune('u', q2_1_1))

  q2_1_1.Add(transition.NewRegexp("[0-9a-f]{4}", q3))

  q3.Add(transition.NewEpsilon(q2))
  q3.Add(transition.NewRune('"', q4))

  str := tomato.NewAutomaton(q0, []*tomato.State{q4})

  items := []struct{
    input string
    recognized bool
  }{
    {"\"\"", true},
    {"\"a\"", true},
    {"\"0\"", true},
    {"\"abcdef\"", true},
    {"abc", false},
    {"\"", false},
    {"\"abc", false},
    {"\"\"\"", false},
    {"\"abc\"\"", false},
  }

  for _, item := range items {
    ok, _ := str.Recognize(item.input)
    if item.recognized && !ok {
      t.Error("Word '"+item.input+"' not recognized")
    }
    if !item.recognized && ok {
      t.Error("Word '"+item.input+"' recognized")
    }
  }
}
