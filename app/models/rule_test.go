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

  rm := New()
  rm.PushBack(r)

  if rm.Check("hello").Content != "hello" {
    t.Log("given key should return matched string configed in rule")
    t.Fail()
  }
}

func TestCheck(t *testing.T) {
  r := &Rule {
    regexp.MustCompile("world"),
    func() (resp *Response){
      resp = NewResponse()
      resp.ToUserName = "FromUserName"
      resp.FromUserName = "ToUserName"
      resp.MsgType = Text
      resp.Content = "world"
      return resp
    },
  }

  rm := New()
  rm.PushBack(r)

  if rm.Check("world").Content != "world" {
    t.Log("given key should return matched string configed in rule")
    t.Fail()
  }
}

func TestNew(t *testing.T) {
  rm := New()
  if rm.rules.Len() != 0 {
    t.Log("rules len should be 0")
    t.Fail()
  }
}

func TestPushBack(t *testing.T) {
  rm := New()

  rm.PushBack(&Rule {
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
