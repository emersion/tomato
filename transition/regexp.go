package transition

import (
  "regexp"

  "github.com/emersion/tomato"
)

type regexpTransition struct {
  re *regexp.Regexp
  target *tomato.State
}

func (t *regexpTransition) Recognize(word string) (*tomato.State, int) {
  loc := t.re.FindStringIndex(word)
  if loc == nil || loc[0] != 0 {
    return nil, 0
  }

  return t.target, loc[1]
}

func NewRegexp(re string, target *tomato.State) tomato.Transition {
  return &regexpTransition{
    re: regexp.MustCompile("^"+re),
    target: target,
  }
}
