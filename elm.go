package elm

//#include <Elementary.h>
import "C"

func NewWin(parent EvasObject, name string, type_ C.Elm_Win_Type) *Win {
  var c_parent *C.Evas_Object
  if parent != nil { c_parent = parent.GetObj() }
  c_name := C.CString(name)
  return &Win{
    obj: C.elm_win_add(c_parent, c_name, type_),
  }
}

func (self *Win) Show() {
  C.evas_object_show(self.obj)
}
