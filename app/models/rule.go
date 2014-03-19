package models

import(
  "container/list"
  "regexp"
  "reflect"
  "github.com/codegangsta/inject"
)

type Rule struct {
  pattern *regexp.Regexp
  handler interface{}
}

type RuleManager struct {
  rules *list.List
  injector inject.Injector
}

func New() (*RuleManager){
  return &RuleManager {
    list.New(),
    inject.New(),
  }
}

func (rm *RuleManager) PushBack(rule *Rule) (*RuleManager) {
  rm.rules.PushBack(rule)
  return rm
}

func (rm *RuleManager) Check(key string) (*Response){
  for e := rm.rules.Front(); e != nil; e = e.Next() {
    rule := e.Value.(*Rule)
    if rule.pattern.Match([]byte(key)) {
      switch reflect.TypeOf(rule.handler).Kind() {
      case reflect.Func:
        ret, _ := rm.injector.Invoke(rule.handler)
        return ret[0].Interface().(*Response)
      case reflect.String:
        resp := NewResponse()
        resp.ToUserName = "FromUserName"
        resp.FromUserName = "ToUserName"
        resp.MsgType = Text
        resp.Content = rule.handler.(string)
        return resp
      }
    }
  }
  return nil
}
