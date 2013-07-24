package main

import (
  "strings"
)

func (self *Generator) collectEinaFuncs() {
  for _, fun := range self.CFuncs {
    if strings.HasPrefix(fun.Name, "eina_") && strings.HasSuffix(fun.Name, "_new") {
      //p("%s\n", fun.Name)
    }
  }
}

func (self *Generator) generateEinaFuncs() {
}
