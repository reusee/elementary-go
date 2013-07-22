package main

import (
  "fmt"
)

type TypeMapFunc func(name, t string) (mappedName, mappedType string, helperCodes []string)

var TYPE_MAPPINGS = map[string]TypeMapFunc{
  "Evas_Object *": func(name, t string) (string, string, []string) {
    return "c_" + name, "EvasObject", []string{
      fmt.Sprintf("var c_%s *C.Evas_Object", name),
      fmt.Sprintf("if %s != nil { c_%s = %s.GetObj() }", name, name, name),
    }
  },
  "void *": func(name, t string) (string, string, []string) {
    return name, "unsafe.Pointer", nil
  },
  "const char *": func(name, t string) (string, string, []string) {
    return "c_" + name, "string", []string{
      fmt.Sprintf("c_%s := C.CString(%s)", name, name),
    }
  },

  // enum types
  "Elm_Win_Type": func(name, t string) (string, string, []string) {
    return name, "C.Elm_Win_Type", nil
  },
}
