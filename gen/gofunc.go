package main

import (
  "go/token"
  "log"
)

type BridgeFunc struct {
  Receiver string
  Name string
  ParamNames []string
  ParamTypes []string
  ReturnTypes []string
  HelperCodes []string
  CgoFunc string
  CgoArguments []string
  ReturnExpression string
}

func (self *BridgeFunc) Gen() string {
  ret := "func " + self.Name + "("
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
  ret += "  _cgo_return_ := C." + self.CgoFunc + "("
  for i, arg := range self.CgoArguments {
    if i > 0 { ret += ", " }
    ret += arg
  }
  ret += ")\n"
  ret += "  return " + self.ReturnExpression + "\n"
  ret += "}\n"
  return ret
}

func (self *BridgeFunc) ConvertParam(name, t string) {
  if tok := token.Lookup(name); tok.IsKeyword() {
    name = name + "_"
  }
  self.ParamNames = append(self.ParamNames, name)
  mapFunc, ok := TYPE_MAPPINGS[t]
  if !ok { log.Fatalf("no map for type %s", t) }
  mappedName, mappedType, helperCodes := mapFunc(name, t)
  self.CgoArguments = append(self.CgoArguments, mappedName)
  self.ParamTypes = append(self.ParamTypes, mappedType)
  self.HelperCodes = append(self.HelperCodes, helperCodes...)
}
