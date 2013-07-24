package main

type Generator struct {
  FuncInfos [][]string
  EnumInfos [][]string
  TypedefInfos [][]string

  CFuncs []*CFunc
  CEnums map[string]string

  Classes []*Class

  EinaFuncs []*BridgeFunc
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

  generator.collectEinaFuncs()
  generator.generateEinaFuncs()

  generator.stat()
}
