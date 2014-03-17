package models

type Rule struct {
  pattern string
  handler func() string
}

var rules []*Rule

func AppendRule(r *Rule) {
  rules = append(rules, r)
}

func Check(key string) string {
  for _, r := range rules {
    if key == r.pattern {
      return r.handler()
    }
  }
  return ""
}

func length() int {
  return len(rules)
}
