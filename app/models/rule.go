package models

import(
  "container/list"
  "regexp"
)

type Rule struct {
  pattern *regexp.Regexp
  handler func() *Response
}

type RuleManager struct {
  rules *list.List
}

func New() (*RuleManager){
  return &RuleManager {
    list.New(),
  }
}

func (rm *RuleManager) PushBack(rule *Rule) (*RuleManager) {
  rm.rules.PushBack(rule)
  return rm
}

func (rm *RuleManager) Check(key string) (res *Response){
  for e := rm.rules.Front(); e != nil; e = e.Next() {
    if e.Value.(*Rule).pattern.Match([]byte(key)) {
      return e.Value.(*Rule).handler()
    }
  }
  return nil
}
