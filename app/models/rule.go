package models

import(
  "container/list"
  "regexp"
  "reflect"
  "github.com/codegangsta/inject"
)

type Rule struct {
  Pattern *regexp.Regexp
  Handler interface{}
}

type RuleManager struct {
  rules *list.List
  injector inject.Injector
}

func NewRuleManager() (*RuleManager){
  return &RuleManager {
    list.New(),
    inject.New(),
  }
}

func (rm *RuleManager) Push(rule *Rule) (*RuleManager) {
  rm.rules.PushBack(rule)
  return rm
}

func (rm *RuleManager) Check(req *Request) (*Response){
  for e := rm.rules.Front(); e != nil; e = e.Next() {
    rule := e.Value.(*Rule)
    if rule.Pattern.Match([]byte(req.Content)) {
      if res, ok := rule.Handler.(*Response); ok {
        return res
      }

      if content, ok := rule.Handler.(string); ok {
        resp := NewResponse()
        resp.MsgType = Text
        resp.FromUserName = req.ToUserName
        resp.ToUserName = req.FromUserName
        resp.Content = content
        return resp
      }

      if reflect.TypeOf(rule.Handler).Kind() == reflect.Func {
        ret, _ := rm.injector.Invoke(rule.Handler)
        return ret[0].Interface().(*Response)
      }
    }
  }
  //default
  resp := NewResponse()
  resp.MsgType = Text
  resp.Content = ""
  return resp
}
