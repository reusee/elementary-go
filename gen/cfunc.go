package main

import (
)

//XXX
var C_MODULES = []string{
  "e_",
  "ecore_",
  "ecvt_",
  "edje_",
  "eet_",
  "efreet_",
  "eina_",
  "elm_",
  "ethumb_",
  "evas_",
}

type CFunc struct {
  ReturnType string
  Name string
  ParamNames []string
  ParamTypes []string
}
