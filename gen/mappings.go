package main

import (
  "fmt"
)

type ParamMapFunc func(name string) (mappedName, mappedType string, helperCodes []string)

var PARAM_MAPPINGS = map[string]ParamMapFunc{
  "Evas_Object *": func(name string) (string, string, []string) {
    return "_c_" + name, "EvasObjectInterface", []string{
      fmt.Sprintf("var _c_%s *C.Evas_Object", name),
      fmt.Sprintf("if %s != nil { _c_%s = %s.GetObj() }", name, name, name),
    }
  },
  "Eina_Bool": func(name string) (string, string, []string) {
    return "_c_" + name, "bool", []string{
      fmt.Sprintf("_c_%s := (C.Eina_Bool)(0)", name),
      fmt.Sprintf("if %s { _c_%s = (C.Eina_Bool)(1) }", name, name),
    }
  },

  "void *": func(name string) (string, string, []string) {
    return name, "unsafe.Pointer", nil
  },
  "const void *": func(name string) (string, string, []string) {
    return name, "unsafe.Pointer", nil
  },
  "const char *": func(name string) (string, string, []string) {
    return "_c_" + name, "string", []string{
      fmt.Sprintf("_c_%s := C.CString(%s)", name, name),
    }
  },
  "char *": func(name string) (string, string, []string) {
    return "_c_" + name, "string", []string{
      fmt.Sprintf("_c_%s := C.CString(%s)", name, name),
      fmt.Sprintf("defer C.free(unsafe.Pointer(_c_%s))", name),
    }
  },

  // numeric types
  "short": func(name string) (string, string, []string) {
    return "_c_" + name, "int", []string{
      fmt.Sprintf("_c_%s := C.short(%s)", name, name),
    }
  },
  "unsigned short": func(name string) (string, string, []string) {
    return "_c_" + name, "uint", []string{
      fmt.Sprintf("_c_%s := C.ushort(%s)", name, name),
    }
  },
  "int": func(name string) (string, string, []string) {
    return "_c_" + name, "int", []string{
      fmt.Sprintf("_c_%s := C.int(%s)", name, name),
    }
  },
  "unsigned int": func(name string) (string, string, []string) {
    return "_c_" + name, "uint", []string{
      fmt.Sprintf("_c_%s := C.uint(%s)", name, name),
    }
  },
  "double": func(name string) (string, string, []string) {
    return "_c_" + name, "float64", []string{
      fmt.Sprintf("_c_%s := C.double(%s)", name, name),
    }
  },

}

type ReturnMapFunc func() (mappedType string, helperCodes []string)

var RETURN_MAPPINGS = map[string]ReturnMapFunc{
  "void *": func() (string, []string) {
    return "unsafe.Pointer", []string{
      "_go_return_ := unsafe.Pointer(_cgo_return_)",
    }
  },
  "Eina_Bool": func() (string, []string) {
    return "bool", []string{
      "_go_return_ := _cgo_return_ == (C.Eina_Bool)(1)",
    }
  },
  "const char *": func() (string, []string) {
    return "string", []string{
      "_go_return_ := C.GoString(_cgo_return_)",
    }
  },
  "Evas_Object *": func() (string, []string) {
    return "*EvasObject", []string{
      "_go_return_ := &EvasObject{_cgo_return_}",
    }
  },
  "Evas_Object_Box_Option *": func() (string, []string) {
    return "*C.Evas_Object_Box_Option", []string{
      "_go_return_ := _cgo_return_",
    }
  },
}
