package transition

import (
  "regexp"

  "github.com/emersion/tomato"
)

func Regexp(re string) Func {
  r := regexp.MustCompile("^"+re)
  return func (word string) int {
    loc := r.FindStringIndex(word)
    if loc == nil || loc[0] != 0 {
      return -1
    }
    return loc[1]
  }
}

func NewRegexp(re string, target *tomato.State) tomato.Transition {
  return NewFunc(Regexp(re), target)
}
