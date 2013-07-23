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
  "const char *[]": func(name string) (string, string, []string) {
    return "_c_" + name, "[]string", []string{
      fmt.Sprintf("_c_%s := ConvertStringSliceToC(%s)", name, name),
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

  "size_t": func(name string) (string, string, []string) {
    return "_c_" + name, "uint64", []string{
      fmt.Sprintf("_c_%s := C.size_t(%s)", name, name),
    }
  },

}

type ReturnMapFunc func() (mappedType string, helperCodes []string)

var RETURN_MAPPINGS = map[string]ReturnMapFunc{
  // numerics
  "short": func() (string, []string) {
    return "int", []string{
      "_go_return_ := int(_cgo_return_)",
    }
  },
  "int": func() (string, []string) {
    return "int", []string{
      "_go_return_ := int(_cgo_return_)",
    }
  },
  "unsigned int": func() (string, []string) {
    return "uint", []string{
      "_go_return_ := uint(_cgo_return_)",
    }
  },
  "double": func() (string, []string) {
    return "float64", []string{
      "_go_return_ := float64(_cgo_return_)",
    }
  },

  // void pointer
  "void *": func() (string, []string) {
    return "unsafe.Pointer", []string{
      "_go_return_ := unsafe.Pointer(_cgo_return_)",
    }
  },
  "const void *": func() (string, []string) {
    return "unsafe.Pointer", []string{
      "_go_return_ := unsafe.Pointer(_cgo_return_)",
    }
  },

  // strings XXX maybe some need to be free, some doesn't
  "const char *": func() (string, []string) {
    return "string", []string{
      "_go_return_ := C.GoString(_cgo_return_)",
    }
  },
  "char *": func() (string, []string) {
    return "string", []string{
      "_go_return_ := C.GoString(_cgo_return_)",
      "C.free(unsafe.Pointer(_cgo_return_))",
    }
  },

  // eina types
  "Eina_Bool": func() (string, []string) {
    return "bool", []string{
      "_go_return_ := _cgo_return_ == (C.Eina_Bool)(1)",
    }
  },

  // evas types
  "Evas_Object *": func() (string, []string) {
    return "*EvasObject", []string{
      "_go_return_ := &EvasObject{_cgo_return_}",
    }
  },
}

type ReturnParamMapFunc func(name string) (mappedType, cType, returnExprs string)

var RETURN_PARAM_MAPPINGS = map[string]ReturnParamMapFunc{
  "unsigned short *": func(name string) (string, string, string) {
    return "uint", "C.ushort", fmt.Sprintf("uint(_c_%s_)", name)
  },
  "double *": func(name string) (string, string, string) {
    return "float64", "C.double", fmt.Sprintf("float64(_c_%s_)", name)
  },
  "int *": func(name string) (string, string, string) {
    return "int", "C.int", fmt.Sprintf("int(_c_%s_)", name)
  },
  "unsigned int *": func(name string) (string, string, string) {
    return "uint", "C.uint", fmt.Sprintf("uint(_c_%s_)", name)
  },
  "const char **": func(name string) (string, string, string) {
    return "string", "*C.char", fmt.Sprintf("C.GoString(_c_%s_)", name)
  },
  "char **": func(name string) (string, string, string) {
    return "string", "*C.char", fmt.Sprintf("ConvertAndFreeCString(_c_%s_)", name)
  },
}
