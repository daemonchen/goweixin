package models

import (
	"testing"
)

func TestAppendRule(t *testing.T) {
  r := &Rule {
    "world",
    func() string{
      return "world"
    },
  }

  oldLen := length()

  AppendRule(r)

  if length() != oldLen+1 {
    t.Log("rule length should increase 1")
    t.Fail()
  }
}

func TestCheck(t *testing.T) {
  r := &Rule {
    "world",
    func() string{
      return "world"
    },
  }

  AppendRule(r)

  if Check("world") != "world" {
    t.Log("given key should return matched string configed in rule")
    t.Fail()
  }
}
