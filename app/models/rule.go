package models

import(
  "container/list"
)

type Rule struct {
  pattern string
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
    if key == e.Value.(*Rule).pattern {
      return e.Value.(*Rule).handler()
    }
  }
  return nil
}
