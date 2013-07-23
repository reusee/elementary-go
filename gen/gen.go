package main

import (
  "fmt"
)

func init() {
  fmt.Printf("")
}

type Generator struct {
  FuncInfos [][]string
  EnumInfos [][]string
  TypedefInfos [][]string

  CFuncs []*CFunc
  CEnums map[string]string

  Classes []*Class
}

func main() {
  generator := new(Generator)
  generator.collectHeaderInfo()
  generator.collectCFuncs()
  generator.collectEnums()
  generator.processTypedefs()

  generator.generateEnums()

  generator.collectClasses()
  generator.generateClasses()

  generator.collectGeneralMethods()
  generator.collectClassMethods()
  generator.generateGeneralMethods()

  generator.stat()
}
