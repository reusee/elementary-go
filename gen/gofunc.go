package main

import (
  "go/token"
  "log"
  "strings"
  "fmt"
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
  CgoHasReturn bool
  ReturnExprs []string
  CFunc *CFunc
}

func (self *BridgeFunc) ConvertParam(name, t string) {
  if tok := token.Lookup(name); tok.IsKeyword() {
    name = name + "_"
  }
  self.ParamNames = append(self.ParamNames, name)
  mapFunc, ok := PARAM_MAPPINGS[t]
  if !ok {
    mapFunc = tryDirectMapParam(t)
    if mapFunc == nil { log.Fatalf("no map for param type %s: %s", t, self.CFunc.Name) }
  }
  mappedName, mappedType, helperCodes := mapFunc(name)
  self.CgoArguments = append(self.CgoArguments, mappedName)
  self.ParamTypes = append(self.ParamTypes, mappedType)
  self.HelperCodes = append(self.HelperCodes, helperCodes...)
}

func (self *BridgeFunc) ConvertReturnType(t string) {
  mapFunc, ok := RETURN_MAPPINGS[t]
  if !ok {
    mapFunc = tryDirectMapReturn(t)
    if mapFunc == nil { log.Fatalf("no map for return type %s: %s", t, self.CFunc.Name) }
  }
  mappedType, helperCodes := mapFunc()
  self.ReturnTypes = append(self.ReturnTypes, mappedType)
  self.HelperCodesAfterCgo = append(self.HelperCodesAfterCgo, helperCodes...)
  self.ReturnExprs = append(self.ReturnExprs, "_go_return_")
}

func (self *BridgeFunc) ConvertReturnParam(name, t string) {
  goType, cType, returnExpr := RETURN_PARAM_MAPPINGS[t](name)
  self.ReturnTypes = append(self.ReturnTypes, goType)
  cVar := fmt.Sprintf("_c_%s_", name)
  self.HelperCodes = append(self.HelperCodes, fmt.Sprintf("var %s %s", cVar, cType))
  self.CgoArguments = append(self.CgoArguments, "&" + cVar)
  self.ReturnExprs = append(self.ReturnExprs, returnExpr)
}

func tryDirectMapParam(t string) ParamMapFunc {
  goType := convertToCgoType(t)
  if goType == "" { return nil }
  return func(name string) (string, string, []string) {
    return name, goType, nil
  }
}

func tryDirectMapReturn(t string) ReturnMapFunc {
  goType := convertToCgoType(t)
  if goType == "" { return nil }
  return func() (string, []string) {
    return goType, []string{
      "_go_return_ := _cgo_return_",
    }
  }
}

func convertToCgoType(t string) string {
  t = strings.Replace(t, "const ", "", -1)
  inModule := false
  for _, m := range C_MODULES {
    if strings.HasPrefix(strings.ToLower(t), m) {
      inModule = true
      break
    }
  }
  if !inModule { return "" }
  ps := 0
  for strings.HasSuffix(t, "*") {
    ps++
    t = strings.TrimSpace(t[:len(t) - 1])
  }
  return strings.Repeat("*", ps) + "C." + t
}

func (self *BridgeFunc) Gen() string {
  if self.CFunc.ReturnType != "void" {
    self.CgoHasReturn = true
  }
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
  if self.CgoHasReturn {
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
  if len(self.ReturnExprs) != 0 {
    ret += "  return " + strings.Join(self.ReturnExprs, ", ") + "\n"
  }
  ret += "}\n"
  return ret
}
