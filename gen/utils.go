package main

import (
  "fmt"
  "sort"
)

func p(format string, args ...interface{}) {
  fmt.Printf(format, args...)
}

// set of string

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

// sort map[string]int

type sortStrIntMap struct {
  m map[string]int
  s []string
}

func (sm *sortStrIntMap) Len() int {
  return len(sm.m)
}

func (sm *sortStrIntMap) Less(i, j int) bool {
  return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortStrIntMap) Swap(i, j int) {
  sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func SortedKeysOfStrIntMap(m map[string]int) []string {
  sm := new(sortStrIntMap)
  sm.m = m
  sm.s = make([]string, len(m))
  i := 0
  for key, _ := range m {
    sm.s[i] = key
    i++
  }
  sort.Sort(sm)
  return sm.s
}
