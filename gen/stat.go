package main

import (
  "fmt"
)

func (self *Generator) stat() {
  n := 0
  for _, fun := range self.CFuncs {
    if fun.Exported { n++ }
  }
  fmt.Printf("exported %d / %d functions\n", n, len(self.CFuncs))
}
