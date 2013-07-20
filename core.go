package elm

/*
#include <Elementary.h>
#include <Emotion.h>
#include <Eio.h>
#include <stdlib.h>
#cgo pkg-config: elementary emotion eio
*/
import "C"
import (
  "os"
  "unsafe"
  "reflect"
)

func init() {
  cargv := make([]*C.char, len(os.Args))
  for i, arg := range os.Args {
    cstr := C.CString(arg)
    defer C.free(unsafe.Pointer(cstr))
    cargv[i] = cstr
  }
  header := (*reflect.SliceHeader)(unsafe.Pointer(&cargv))
  C.elm_init(C.int(len(os.Args)), (**C.char)(unsafe.Pointer(header.Data)))
}

func Run() {
  C.elm_run()
}

func ShutDown() int {
  return int(C.elm_shutdown())
}

func Exit() {
  C.elm_exit()
}
