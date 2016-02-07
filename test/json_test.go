package test

import (
  "testing"

  "github.com/emersion/tomato"
  "github.com/emersion/tomato/transition"
)

func jsonNumber() *tomato.Automaton {
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

  q0.AddFunc(transition.Rune('-'), q1)
  q0.AddFunc(transition.Epsilon(), q1)

  q1.AddFunc(transition.Rune('0'), q2)
  q1.AddFunc(transition.Regexp("[1-9]"), q1_1)

  q1_1.AddFunc(transition.Regexp("[0-9]"), q1_1)
  q1_1.AddFunc(transition.Epsilon(), q2)

  q2.AddFunc(transition.Epsilon(), q5)
  q2.AddFunc(transition.Rune('.'), q3)

  q3.AddFunc(transition.Regexp("[0-9]"), q4)

  q4.AddFunc(transition.Regexp("[0-9]"), q4)
  q4.AddFunc(transition.Epsilon(), q5)

  q5.AddFunc(transition.Regexp("[eE]"), q5_1)
  q5.AddFunc(transition.Epsilon(), q6)

  q5_1.AddFunc(transition.Regexp("[+-]?"), q5_2)

  q5_2.AddFunc(transition.Regexp("[0-9]"), q5_3)

  q5_3.AddFunc(transition.Regexp("[0-9]"), q5_3)
  q5_3.AddFunc(transition.Epsilon(), q6)

  return tomato.NewAutomaton(q0, []*tomato.State{q6})
}

func jsonString() *tomato.Automaton {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q2 := tomato.NewState()
  q2_1 := tomato.NewState()
  q2_1_1 := tomato.NewState()
  q3 := tomato.NewState()
  q4 := tomato.NewState()

  q0.AddFunc(transition.Rune('"'), q1)

  q1.AddFunc(transition.Epsilon(), q2)
  q1.AddFunc(transition.Epsilon(), q3)

  q2.AddFunc(transition.Regexp("[^\\\"]+"), q3)
  q2.AddFunc(transition.Rune('\\'), q2_1)

  q2_1.AddFunc(transition.Regexp("[\"\\/bfnrt]"), q3)
  q2_1.AddFunc(transition.Rune('u'), q2_1_1)

  q2_1_1.AddFunc(transition.Regexp("[0-9a-f]{4}"), q3)

  q3.AddFunc(transition.Epsilon(), q2)
  q3.AddFunc(transition.Rune('"'), q4)

  return tomato.NewAutomaton(q0, []*tomato.State{q4})
}

func jsonValue(str, number, array, object *tomato.Automaton) *tomato.Automaton {
  q0 := tomato.NewState()
  q1 := tomato.NewState()

  q0.Add(transition.NewAutomaton(str, []*tomato.State{q1}))
  q0.Add(transition.NewAutomaton(number, []*tomato.State{q1}))
  q0.Add(transition.NewAutomaton(object, []*tomato.State{q1}))
  q0.Add(transition.NewAutomaton(array, []*tomato.State{q1}))
  q0.AddFunc(transition.String("true"), q1)
  q0.AddFunc(transition.String("false"), q1)
  q0.AddFunc(transition.String("null"), q1)

  return tomato.NewAutomaton(q0, []*tomato.State{q1})
}

func jsonArray(value *tomato.Automaton) *tomato.Automaton {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q2 := tomato.NewState()
  q3 := tomato.NewState()
  q4 := tomato.NewState()
  q5 := tomato.NewState()

  q0.AddFunc(transition.Rune('['), q1)

  q1.AddFunc(transition.Epsilon(), q2)
  q1.AddFunc(transition.Epsilon(), q4)

  q2.Add(transition.NewAutomaton(value, []*tomato.State{q3}))

  q3.AddFunc(transition.Rune(','), q2)
  q3.AddFunc(transition.Epsilon(), q4)

  q4.AddFunc(transition.Rune(']'), q5)

  return tomato.NewAutomaton(q0, []*tomato.State{q4})
}

func jsonObject(str, value *tomato.Automaton) *tomato.Automaton {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q2 := tomato.NewState()
  q3 := tomato.NewState()
  q4 := tomato.NewState()
  q5 := tomato.NewState()
  q6 := tomato.NewState()
  q7 := tomato.NewState()

  q0.AddFunc(transition.Rune('{'), q1)

  q1.AddFunc(transition.Epsilon(), q2)
  q1.AddFunc(transition.Epsilon(), q6)

  q2.Add(transition.NewAutomaton(str, []*tomato.State{q3}))

  q3.AddFunc(transition.Rune(':'), q4)

  q4.Add(transition.NewAutomaton(value, []*tomato.State{q5}))

  q5.AddFunc(transition.Rune(','), q2)
  q5.AddFunc(transition.Epsilon(), q6)

  q6.AddFunc(transition.Rune('}'), q7)

  return tomato.NewAutomaton(q0, []*tomato.State{q7})
}

func TestJsonNumber(t *testing.T) {
  number := jsonNumber()

  items := []testCase{
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

  testAll(t, number, items)
}

func TestJsonString(t *testing.T) {
  str := jsonString()

  items := []testCase{
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

  testAll(t, str, items)
}
