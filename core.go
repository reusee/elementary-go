package elm

/*
#include <Elementary.h>
#include <Emotion.h>
#include <Eio.h>
#include <stdlib.h>

extern void CallCb(void*, void*);
static void call_cb(void *cb, Evas_Object *obj, void *event_info) {
  CallCb(cb, event_info);
}
void hook(Evas_Object *obj, const char *ev, void *data) {
  evas_object_smart_callback_add(obj, ev, call_cb, data);
}

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

func ConvertAndFreeCString(cstr *C.char) string {
  goStr := C.GoString(cstr)
  C.free(unsafe.Pointer(cstr))
  return goStr
}

func ConvertStringSliceToC(ss []string) **C.char {
  cstrs := make([]*C.char, len(ss))
  for i := 0; i < len(ss); i++ {
    cstrs[i] = C.CString(ss[i])
  }
  return &cstrs[0]
}

type EvInfo struct {
  CInfo unsafe.Pointer
}

func fromCEvInfo(info unsafe.Pointer) *EvInfo {
  return &EvInfo{info}
}

type Callback func(*EvInfo)

func Connect(target EvasObjectInterface, ev string, cb Callback) {
  evStr := C.CString(ev)
  C.hook(target.GetObj(), evStr, unsafe.Pointer(&cb))
  C.free(unsafe.Pointer(evStr))
}

var (
  EV_CLICKED = C.CString("clicked")
)

func (self *Button) OnClicked(cb Callback) {
  C.hook(self.obj, EV_CLICKED, unsafe.Pointer(&cb))
}
