package main

import (
  "os"
)

func (self *Generator) stat() {
  notExported, err := os.Create("not_exported")
  moduleStat := make(map[string]int)
  if err != nil { panic(err) }
  defer notExported.Close()
  n := 0
  for _, fun := range self.CFuncs {
    if fun.Exported {
      n++
    } else {
      notExported.Write([]byte(fun.Name))
      notExported.Write([]byte("\n"))
      moduleStat[fun.Module]++
    }
  }
  p("exported %d / %d functions\n", n, len(self.CFuncs))
  p("not expoted:\n")
  for _, k := range SortedKeysOfStrIntMap(moduleStat) {
    p("%s %d\n", k, moduleStat[k])
  }
}
