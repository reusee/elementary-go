package main

import (
  "io/ioutil"
  "log"
  "strings"
  "fmt"
)

func main() {
  infoBs, err := ioutil.ReadFile("header_info")
  if err != nil { log.Fatal(err) }
  cFuncs := make([]CFunc, 0)
  for _, line := range strings.Split(string(infoBs), "\n") {
    lineSp := strings.Split(line, "|")
    if lineSp[0] == "func" { // process function
      cfunc := processCFunc(lineSp)
      inModule := false
      for _, m := range C_MODULES {
        if strings.HasPrefix(cfunc.Name, m) {
          inModule = true
          break
        }
      }
      if !inModule { continue }
      fmt.Printf("%s\n", cfunc.Name)
      cFuncs = append(cFuncs, cfunc)
    }
  }
}

func processCFunc(lineSp []string) CFunc {
  name := lineSp[1]
  returnType := lineSp[2]
  paramNames := make([]string, 0)
  paramTypes := make([]string, 0)
  for _, param := range lineSp[3:] {
    paramSp := strings.Split(param, "@")
    if len(paramSp) == 1 {
      paramNames = append(paramNames, "")
      paramTypes = append(paramTypes, paramSp[0])
    } else {
      paramNames = append(paramNames, paramSp[1])
      paramTypes = append(paramTypes, paramSp[0])
    }
  }
  cfunc := CFunc{
    Name: name,
    ReturnType: returnType,
    ParamNames: paramNames,
    ParamTypes: paramTypes,
  }
  return cfunc
}
