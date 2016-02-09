package test

import (
  "testing"

  "github.com/emersion/tomato"
  "github.com/emersion/tomato/transition"
  "github.com/emersion/tomato/wrapper"
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

  return tomato.NewAutomaton(q0, []*tomato.State{q4})
}

func jsonArray(value *tomato.Automaton) *tomato.Automaton {
  q0 := tomato.NewState()
  q1 := tomato.NewState()
  q2 := tomato.NewState()
  q3 := tomato.NewState()
  q4 := tomato.NewState()
  q5 := tomato.NewState()

  q0.Add(transition.NewRune('[', q1))

  q1.Add(transition.NewEpsilon(q2))
  q1.Add(transition.NewEpsilon(q4))

  q2.Add(transition.NewAutomaton(value, []*tomato.State{q3}))

  q3.Add(transition.NewRune(',', q2))
  q3.Add(transition.NewEpsilon(q4))

  q4.Add(transition.NewRune(']', q5))

  return tomato.NewAutomaton(q0, []*tomato.State{q5})
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

  q0.Add(transition.NewRune('{', q1))

  q1.Add(transition.NewEpsilon(q2))
  q1.Add(transition.NewEpsilon(q6))

  q2.Add(transition.NewAutomaton(str, []*tomato.State{q3}))

  q3.Add(transition.NewRune(':', q4))

  q4.Add(transition.NewAutomaton(value, []*tomato.State{q5}))

  q5.Add(transition.NewRune(',', q2))
  q5.Add(transition.NewEpsilon(q6))

  q6.Add(transition.NewRune('}', q7))

  return tomato.NewAutomaton(q0, []*tomato.State{q7})
}

func jsonValue(number, str, array, object *tomato.Automaton) *tomato.Automaton {
  q0 := tomato.NewState()
  q1 := tomato.NewState()

  q0.Add(transition.NewAutomaton(str, []*tomato.State{q1}))
  q0.Add(transition.NewAutomaton(number, []*tomato.State{q1}))
  q0.Add(transition.NewAutomaton(object, []*tomato.State{q1}))
  q0.Add(transition.NewAutomaton(array, []*tomato.State{q1}))
  q0.Add(transition.NewString("true", q1))
  q0.Add(transition.NewString("false", q1))
  q0.Add(transition.NewString("null", q1))

  return tomato.NewAutomaton(q0, []*tomato.State{q1})
}

func json() (number, str, array, object, value *tomato.Automaton) {
  valueWrapper, setValue := wrapper.New(1)

  number = jsonNumber()
  str = jsonString()
  array = jsonArray(valueWrapper)
  object = jsonObject(str, valueWrapper)
  value = jsonValue(number, str, array, object)

  setValue(value)

  return
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

func TestJsonValue(t *testing.T) {
  _, _, array, object, _ := json()

  testAll(t, array, []testCase{
    {"[]", true},
    {"[\"a\"]", true},
    {"[\"a\",\"b\"]", true},
    {"[\"a\",\"b\",\"c\"]", true},
    {"[0,1,2]", true},
    {"[", false},
    {"[][", false},
    {"[a]", false},
    {"[\"a]", false},
    {"[\"a\",]", false},
    {"[\"a\",\"b\",]", false},
    {"[\"a\"\"b\"]", false},
  })

  testAll(t, object, []testCase{
    {"{}", true},
    {"{\"a\":0}", true},
    {"{\"a\":0,\"b\":1}", true},
    {"{\"a\":\"b\",\"c\":\"d\"}", true},
    {"{\"a\":\"b\",\"c\":4}", true},
    {"{\"a\":[0,1],\"c\":[2,3]}", true},
    {"{\"a\":{\"a\":true,\"b\":false},\"c\":null}", true},
    {"{", false},
    {"}", false},
    {"{a}", false},
    {"{\"a:0}", false},
    {"{\"a\"0}", false},
    {"{\"a\"}", false},
    {"{\"a\":0,}", false},
    {"{\"a\":0,\"b\":1,}", false},
  })
}
