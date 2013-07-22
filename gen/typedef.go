package main

import (
  "strings"
)

func processTypedef(lineSp []string) {
  t1 := lineSp[1]
  t2 := lineSp[2]
  t3 := lineSp[3]
  numericTypes := map[string]bool {
    "int": true,
    "unsigned long long": true,
  }
  if _, has := numericTypes[t2]; strings.ToLower(t1) != t1 && has { // typedef ___ [numeric];
    PARAM_MAPPINGS[t1] = func(name, t string) (string, string, []string) {
      return name, "C." + t1, nil
    }
  } else if strings.HasPrefix(t2, "enum ") && t1 == t3 { // typedef enum;
    PARAM_MAPPINGS[t1] = func(name, t string) (string, string, []string) {
      return name, "C." + t1, nil
    }
  }
}
