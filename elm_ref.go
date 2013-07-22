package elm

//#include <Elementary.h>
import "C"

func (self *Win) Show() {
  C.evas_object_show(self.obj)
}
