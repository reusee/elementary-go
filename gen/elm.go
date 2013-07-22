package main

import (
  "strings"
  "fmt"
  "os"
)

//TODO
//XXX
var SPECIAL_CONSTRUCT_FUNCS = map[string]bool {
  "elm_win_util_standard_add": true,
}

type Class struct {
  Name string
  CConstructFunc string
}

func genElmClasses(funcs []CFunc) {
  // find all constructors
  classes := make([]Class, 0)
  for _, fun := range funcs {
    if strings.HasPrefix(fun.Name, "elm_") && strings.HasSuffix(fun.Name, "_add") && fun.ReturnType == "Evas_Object *" {
      if _, has := SPECIAL_CONSTRUCT_FUNCS[fun.Name]; has { continue }
      fmt.Printf("%s %v %v\n", fun.Name, fun.ParamTypes, fun.ParamNames)
      className := convertToClassName(fun.Name)
      class := Class{
        Name: className,
        CConstructFunc: fun.Name,
      }
      classes = append(classes, class)
    }
  }

  genClasses(classes)
}

func convertToClassName(name string) string {
  name = name[len("elm_") : len(name) - len("_add")]
  seg := strings.Split(name, "_")
  for i, s := range seg {
    seg[i] = strings.Title(s)
  }
  name = strings.Join(seg, "")
  fmt.Printf("%s\n", name)
  return name
}

func genClasses(classes []Class) {
  outputFile, err := os.Create("../elm_widgets.go")
  if err != nil { panic(err) }
  defer outputFile.Close()
  outputFile.Write([]byte(`package elm

//#include <Elementary.h>
import "C"

type EvasObject interface {
  GetObj() *C.Evas_Object
}
`))
  for _, class := range classes {
    fmt.Fprintf(outputFile, `
type %s struct { obj *C.Evas_Object }
func (self *%s) GetObj() *C.Evas_Object { return self.obj }
`, class.Name, class.Name)
  }
}
