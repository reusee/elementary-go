package main

import (
  "io/ioutil"
  "log"
  "strings"
  "fmt"
)

func init() {
  fmt.Printf("")
}

func main() {
  infoBs, err := ioutil.ReadFile("header_info")
  if err != nil { log.Fatal(err) }

  cFuncs := make([]CFunc, 0)
  cEnums := make(map[string]string)

  for _, line := range strings.Split(string(infoBs), "\n") {
    lineSp := strings.Split(line, "|")
    // process function
    if lineSp[0] == "func" {
      cfunc := processCFunc(lineSp)
      inModule := false
      for _, m := range C_MODULES {
        if strings.HasPrefix(cfunc.Name, m) {
          inModule = true
          break
        }
      }
      if !inModule { continue }
      cFuncs = append(cFuncs, cfunc)
    // process enum
    } else if lineSp[0] == "enum" {
      name := lineSp[1]
      enumloop: for _, m := range C_MODULES {
        m = strings.ToUpper(m)
        if strings.HasPrefix(name, m) {
          name = name[len(m):]
          if am, has := cEnums[name]; has {
            if preferM, has := PREFER_ENUM[name]; has {
              cEnums[name] = preferM
            } else {
              log.Fatalf("enum conflict: %s %s %s, add entry to PREFER_ENUM to resolve\n", name, m, am)
            }
          }
          cEnums[name] = m
          break enumloop
        }
      }
    }
  }

  genEnums(cEnums)
  genElmClasses(cFuncs)
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
