package main

import (
  "strings"
  "fmt"
  "os"
)

type Class struct {
  Name string
  CConstructor *CFunc
  GoConstructFunc *BridgeFunc
  Methods []*BridgeFunc
}

func (self *Generator) collectClasses() {
  for _, fun := range self.CFuncs {
    if strings.HasPrefix(fun.Name, "elm_") && strings.HasSuffix(fun.Name, "_add") && fun.ReturnType == "Evas_Object *" {
      if DISCARD_CONSTRUCT_FUNCS.Has(fun.Name) { continue }
      className := convertToClassName(fun.Name)
      class := &Class{
        Name: className,
        CConstructor: fun,
        GoConstructFunc: makeGoConstructFunc(className, fun),
      }
      self.Classes = append(self.Classes, class)
      fun.Exported = true
    }
  }
}

func (self *Generator) generateClasses() {
  outputFile, err := os.Create("../class.go")
  if err != nil { panic(err) }
  defer outputFile.Close()
  outputFile.Write([]byte(`package elm // generated by gen/class.go

//#include <Elementary.h>
import "C"
import (
  "unsafe"
)

type EvasObject struct {
  obj *C.Evas_Object
}
func (self *EvasObject) GetObj() *C.Evas_Object {
  return self.obj
}

type EvasObjectInterface interface {
  GetObj() *C.Evas_Object
}
`))
  for _, class := range self.Classes {
    fmt.Fprintf(outputFile, `
type %s struct { obj *C.Evas_Object }
func (self *%s) GetObj() *C.Evas_Object { return self.obj }
`, class.Name, class.Name)
    outputFile.Write([]byte("\n"))
    outputFile.Write([]byte(class.GoConstructFunc.Gen()))
  }
}

func convertToClassName(name string) string {
  name = name[len("elm_") : len(name) - len("_add")]
  seg := strings.Split(name, "_")
  for i, s := range seg {
    seg[i] = strings.Title(s)
  }
  name = strings.Join(seg, "")
  return name
}

func makeGoConstructFunc(className string, fun *CFunc) *BridgeFunc {
  gofunc := new(BridgeFunc)
  gofunc.CFunc = fun
  gofunc.Name = "New" + className
  for i, name := range fun.ParamNames {
    gofunc.ConvertParam(name, fun.ParamTypes[i])
  }
  gofunc.ReturnTypes = []string{"*" + className}
  gofunc.CgoFunc = fun.Name
  gofunc.ReturnExprs = append(gofunc.ReturnExprs, "&" + className + "{obj: _cgo_return_}")
  return gofunc
}