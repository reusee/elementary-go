package main

import (
  "log"
  "io/ioutil"
  "bytes"
  "strings"
)

func (self *Generator) collectHeaderInfo() {
  infoBs, err := ioutil.ReadFile("header_info")
  if err != nil { log.Fatal(err) }
  for _, lineBs := range bytes.Split(infoBs, []byte("\n")) {
    if len(lineBs) == 0 { continue }
    lineSp := strings.Split(string(lineBs), "|")
    switch lineSp[0] {
    case "func":
      self.FuncInfos = append(self.FuncInfos, lineSp[1:])
    case "enum":
      self.EnumInfos = append(self.EnumInfos, lineSp[1:])
    case "typedef":
      self.TypedefInfos = append(self.TypedefInfos, lineSp[1:])
    case "func:variadic":
      // discard
    default:
      log.Fatal("not handle header info type ", lineSp[0])
    }
  }
}
