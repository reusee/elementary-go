package main

import (
  "go/token"
  "log"
  "strings"
)

type BridgeFunc struct {
  Receiver string
  Name string
  ParamNames []string
  ParamTypes []string
  ReturnTypes []string
  HelperCodes []string
  HelperCodesAfterCgo []string
  CgoFunc string
  CgoArguments []string
  ReturnExpression string
  CFunc *CFunc
}

func (self *BridgeFunc) ConvertParam(name, t string) {
  if tok := token.Lookup(name); tok.IsKeyword() {
    name = name + "_"
  }
  self.ParamNames = append(self.ParamNames, name)
  mapFunc, ok := PARAM_MAPPINGS[t]
  if !ok {
    // try to map it directory
    mapFunc = tryDirectMap(t)
    if mapFunc == nil { log.Fatalf("no map for param type %s", t) }
  }
  mappedName, mappedType, helperCodes := mapFunc(name)
  self.CgoArguments = append(self.CgoArguments, mappedName)
  self.ParamTypes = append(self.ParamTypes, mappedType)
  self.HelperCodes = append(self.HelperCodes, helperCodes...)
}

func (self *BridgeFunc) ConvertReturnType(t string) {
  mapFunc, ok := RETURN_MAPPINGS[t]
  if !ok { log.Fatalf("no map for return type %s", t) }
  mappedType, helperCodes := mapFunc()
  self.ReturnTypes = append(self.ReturnTypes, mappedType)
  self.HelperCodesAfterCgo = append(self.HelperCodesAfterCgo, helperCodes...)
}

func tryDirectMap(t string) ParamMapFunc {
  t = strings.Replace(t, "const ", "", -1)
  inModule := false
  for _, m := range C_MODULES {
    if strings.HasPrefix(strings.ToLower(t), m) {
      inModule = true
      break
    }
  }
  if !inModule { return nil }
  ps := 0
  for strings.HasSuffix(t, "*") {
    ps++
    t = strings.TrimSpace(t[:len(t) - 1])
  }
  goType := strings.Repeat("*", ps) + "C." + t
  return func(name string) (string, string, []string) {
    return name, goType, nil
  }
}

func (self *BridgeFunc) Gen() string {
  ret := "func"
  if self.Receiver != "" {
    ret += " (self " + self.Receiver + ")"
  }
  ret += " " + self.Name + "("
  for i, name := range self.ParamNames {
    if i > 0 { ret += ", " }
    ret += name + " " + self.ParamTypes[i]
  }
  ret += ") ("
  for i, t := range self.ReturnTypes {
    if i > 0 { ret += ", " }
    ret += t
  }
  ret += ") {\n"
  for _, code := range self.HelperCodes {
    ret += "  " + code + "\n"
  }
  ret += "  "
  if len(self.ReturnTypes) != 0 {
    ret += "_cgo_return_ := "
  }
  ret += "C." + self.CgoFunc + "("
  for i, arg := range self.CgoArguments {
    if i > 0 { ret += ", " }
    ret += arg
  }
  ret += ")\n"
  for _, code := range self.HelperCodesAfterCgo {
    ret += "  " + code + "\n"
  }
  if len(self.ReturnTypes) != 0 {
    ret += "  return " + self.ReturnExpression + "\n"
  }
  ret += "}\n"
  return ret
}
