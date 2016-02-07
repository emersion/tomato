package transition

import (
  "regexp"

  "github.com/emersion/tomato"
)

func Regexp(reStr string) tomato.TransitionFunc {
  re := regexp.MustCompile("^"+reStr)
  return func (word string) int {
    loc := re.FindStringIndex(word)
    if loc == nil || loc[0] != 0 {
      return -1
    }
    return loc[1]
  }
}
