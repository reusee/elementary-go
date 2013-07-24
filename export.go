package elm

import "C"
import (
  "unsafe"
)

//export CallCb
func CallCb(cbp unsafe.Pointer, ev unsafe.Pointer) {
  (*(*Callback)(cbp))(fromCEvInfo(ev))
}
