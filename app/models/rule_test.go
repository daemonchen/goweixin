package models

import (
	"testing"
)

func TestCheck(t *testing.T) {
  r := &Rule {
    "world",
    func() string{
      return "world"
    },
  }

  rm := New()
  rm.PushBack(r)

  if rm.Check("world") != "world" {
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
    "world",
    func() string{
      return "world"
    },
  })
  if rm.rules.Len() != 1 {
    t.Log("rules len should be 0")
    t.Fail()
  }
}
