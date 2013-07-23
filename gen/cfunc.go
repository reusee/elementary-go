package main

import (
  "strings"
)

type CFunc struct {
  ReturnType string
  Name string
  ParamNames []string
  ParamTypes []string
  Exported bool
  Module string
}

func (self *Generator) collectCFuncs() {
  for _, info := range self.FuncInfos {
    cfunc := makeCFunc(info)
    inModule := false
    for _, m := range C_MODULES {
      if strings.HasPrefix(cfunc.Name, m) {
        cfunc.Module = m
        inModule = true
        break
      }
    }
    if !inModule {
      continue
    }
    self.CFuncs = append(self.CFuncs, cfunc)
  }
}

func makeCFunc(info []string) *CFunc {
  name := info[0]
  returnType := info[1]
  paramNames := make([]string, 0)
  paramTypes := make([]string, 0)
  for _, param := range info[2:] {
    paramSp := strings.Split(param, "@")
    if len(paramSp) == 1 {
      paramNames = append(paramNames, "")
      paramTypes = append(paramTypes, paramSp[0])
    } else {
      paramNames = append(paramNames, paramSp[1])
      paramTypes = append(paramTypes, paramSp[0])
    }
  }
  cfunc := &CFunc{
    Name: name,
    ReturnType: returnType,
    ParamNames: paramNames,
    ParamTypes: paramTypes,
  }
  return cfunc
}
