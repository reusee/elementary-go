package main

import (
  "strings"
)

func (self *Generator) processTypedefs() {
  for _, info := range self.TypedefInfos {
    processTypedef(info)
  }
}

func processTypedef(info []string) {
  t1 := info[0]
  t2 := info[1]
  t3 := info[2]
  _ = t3
  numericTypes := map[string]bool {
    "int": true,
    "unsigned long long": true,
  }
  if _, has := numericTypes[t2]; strings.ToLower(t1) != t1 && has { // typedef ___ [numeric];
    PARAM_MAPPINGS[t1] = func(name string) (string, string, []string) {
      return name, "C." + t1, nil
    }
  } else if strings.HasPrefix(t2, "enum ") { // typedef enum;
    PARAM_MAPPINGS[t1] = func(name string) (string, string, []string) {
      return name, "C." + t1, nil
    }
    RETURN_MAPPINGS[t1] = func() (string, []string) {
      return "C." + t1, []string{
        "_go_return_ := _cgo_return_",
      }
    }
  }
}
