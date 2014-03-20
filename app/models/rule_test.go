package models

import (
	"testing"
  "regexp"
)

func TestStringHandler(t *testing.T) {
  r := &Rule {
    regexp.MustCompile("hello"),
    "hello",
  }

  rm := NewRuleManager()
  rm.Push(r)

  req := &Request{}
  req.MsgType = "Text"
  req.Content = "hello"
  req.FromUserName = "FromUserName"
  req.ToUserName = "ToUserName"

  if rm.Check(req).Content != "hello" {
    t.Log("given key should return matched string configed in rule")
    t.Fail()
  }
}

func TestResponseHandler(t *testing.T) {
  resp := NewResponse()
  resp.ToUserName = "FromUserName"
  resp.FromUserName = "ToUserName"
  resp.MsgType = Text
  resp.Content = "world"
  r := &Rule {
    regexp.MustCompile("hello"),
    resp,
  }

  rm := NewRuleManager()
  rm.Push(r)

  req := &Request{}
  req.MsgType = "Text"
  req.Content = "hello"
  req.FromUserName = "FromUserName"
  req.ToUserName = "ToUserName"

  if rm.Check(req).Content != "world" {
    t.Log("given key should return matched string configed in rule")
    t.Fail()
  }
}

func TestCheck(t *testing.T) {
  r := &Rule {
    regexp.MustCompile("hello"),
    func() (resp *Response){
      resp = NewResponse()
      resp.ToUserName = "FromUserName"
      resp.FromUserName = "ToUserName"
      resp.MsgType = Text
      resp.Content = "world"
      return resp
    },
  }

  req := &Request{}
  req.MsgType = "Text"
  req.Content = "hello"
  req.FromUserName = "FromUserName"
  req.ToUserName = "ToUserName"

  rm := NewRuleManager()
  rm.Push(r)

  if rm.Check(req).Content != "world" {
    t.Log("given key should return matched string configed in rule")
    t.Fail()
  }
}

func TestNew(t *testing.T) {
  rm := NewRuleManager()
  if rm.rules.Len() != 0 {
    t.Log("rules len should be 0")
    t.Fail()
  }
}

func TestPush(t *testing.T) {
  rm := NewRuleManager()

  rm.Push(&Rule {
    regexp.MustCompile("world"),
    func() (resp *Response) {
      resp = NewResponse()
      resp.ToUserName = "FromUserName"
      resp.FromUserName = "ToUserName"
      resp.MsgType = Text
      return resp
    },
  })
  if rm.rules.Len() != 1 {
    t.Log("rules len should be 0")
    t.Fail()
  }
}