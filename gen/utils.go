package main

import (
  "fmt"
)

func p(format string, args ...interface{}) {
  fmt.Printf(format, args...)
}

type StringSet struct {
  m map[string]bool
}

func NewStringSet(strs []string) *StringSet {
  self := &StringSet{
    m: make(map[string]bool),
  }
  for _, s := range strs {
    self.m[s] = true
  }
  return self
}

func (self *StringSet) Has(s string) bool {
  _, has := self.m[s]
  return has
}
