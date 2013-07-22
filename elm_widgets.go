package elm

//#include <Elementary.h>
import "C"

type EvasObject interface {
  GetObj() *C.Evas_Object
}

type Icon struct { obj *C.Evas_Object }
func (self *Icon) GetObj() *C.Evas_Object { return self.obj }

type Scroller struct { obj *C.Evas_Object }
func (self *Scroller) GetObj() *C.Evas_Object { return self.obj }

type Entry struct { obj *C.Evas_Object }
func (self *Entry) GetObj() *C.Evas_Object { return self.obj }

type List struct { obj *C.Evas_Object }
func (self *List) GetObj() *C.Evas_Object { return self.obj }

type Ctxpopup struct { obj *C.Evas_Object }
func (self *Ctxpopup) GetObj() *C.Evas_Object { return self.obj }

type Dayselector struct { obj *C.Evas_Object }
func (self *Dayselector) GetObj() *C.Evas_Object { return self.obj }

type FileselectorButton struct { obj *C.Evas_Object }
func (self *FileselectorButton) GetObj() *C.Evas_Object { return self.obj }

type FileselectorEntry struct { obj *C.Evas_Object }
func (self *FileselectorEntry) GetObj() *C.Evas_Object { return self.obj }

type Fileselector struct { obj *C.Evas_Object }
func (self *Fileselector) GetObj() *C.Evas_Object { return self.obj }

type Hoversel struct { obj *C.Evas_Object }
func (self *Hoversel) GetObj() *C.Evas_Object { return self.obj }

type Multibuttonentry struct { obj *C.Evas_Object }
func (self *Multibuttonentry) GetObj() *C.Evas_Object { return self.obj }

type Naviframe struct { obj *C.Evas_Object }
func (self *Naviframe) GetObj() *C.Evas_Object { return self.obj }

type Popup struct { obj *C.Evas_Object }
func (self *Popup) GetObj() *C.Evas_Object { return self.obj }

type Actionslider struct { obj *C.Evas_Object }
func (self *Actionslider) GetObj() *C.Evas_Object { return self.obj }

type Bg struct { obj *C.Evas_Object }
func (self *Bg) GetObj() *C.Evas_Object { return self.obj }

type Box struct { obj *C.Evas_Object }
func (self *Box) GetObj() *C.Evas_Object { return self.obj }

type Bubble struct { obj *C.Evas_Object }
func (self *Bubble) GetObj() *C.Evas_Object { return self.obj }

type Button struct { obj *C.Evas_Object }
func (self *Button) GetObj() *C.Evas_Object { return self.obj }

type Calendar struct { obj *C.Evas_Object }
func (self *Calendar) GetObj() *C.Evas_Object { return self.obj }

type Check struct { obj *C.Evas_Object }
func (self *Check) GetObj() *C.Evas_Object { return self.obj }

type Clock struct { obj *C.Evas_Object }
func (self *Clock) GetObj() *C.Evas_Object { return self.obj }

type Colorselector struct { obj *C.Evas_Object }
func (self *Colorselector) GetObj() *C.Evas_Object { return self.obj }

type Conformant struct { obj *C.Evas_Object }
func (self *Conformant) GetObj() *C.Evas_Object { return self.obj }

type Datetime struct { obj *C.Evas_Object }
func (self *Datetime) GetObj() *C.Evas_Object { return self.obj }

type Diskselector struct { obj *C.Evas_Object }
func (self *Diskselector) GetObj() *C.Evas_Object { return self.obj }

type Flip struct { obj *C.Evas_Object }
func (self *Flip) GetObj() *C.Evas_Object { return self.obj }

type Flipselector struct { obj *C.Evas_Object }
func (self *Flipselector) GetObj() *C.Evas_Object { return self.obj }

type Frame struct { obj *C.Evas_Object }
func (self *Frame) GetObj() *C.Evas_Object { return self.obj }

type Gengrid struct { obj *C.Evas_Object }
func (self *Gengrid) GetObj() *C.Evas_Object { return self.obj }

type Genlist struct { obj *C.Evas_Object }
func (self *Genlist) GetObj() *C.Evas_Object { return self.obj }

type GestureLayer struct { obj *C.Evas_Object }
func (self *GestureLayer) GetObj() *C.Evas_Object { return self.obj }

type Glview struct { obj *C.Evas_Object }
func (self *Glview) GetObj() *C.Evas_Object { return self.obj }

type Grid struct { obj *C.Evas_Object }
func (self *Grid) GetObj() *C.Evas_Object { return self.obj }

type Hover struct { obj *C.Evas_Object }
func (self *Hover) GetObj() *C.Evas_Object { return self.obj }

type Image struct { obj *C.Evas_Object }
func (self *Image) GetObj() *C.Evas_Object { return self.obj }

type Index struct { obj *C.Evas_Object }
func (self *Index) GetObj() *C.Evas_Object { return self.obj }

type WinInwin struct { obj *C.Evas_Object }
func (self *WinInwin) GetObj() *C.Evas_Object { return self.obj }

type Label struct { obj *C.Evas_Object }
func (self *Label) GetObj() *C.Evas_Object { return self.obj }

type Layout struct { obj *C.Evas_Object }
func (self *Layout) GetObj() *C.Evas_Object { return self.obj }

type Map struct { obj *C.Evas_Object }
func (self *Map) GetObj() *C.Evas_Object { return self.obj }

type MapTrack struct { obj *C.Evas_Object }
func (self *MapTrack) GetObj() *C.Evas_Object { return self.obj }

type Mapbuf struct { obj *C.Evas_Object }
func (self *Mapbuf) GetObj() *C.Evas_Object { return self.obj }

type Menu struct { obj *C.Evas_Object }
func (self *Menu) GetObj() *C.Evas_Object { return self.obj }

type Notify struct { obj *C.Evas_Object }
func (self *Notify) GetObj() *C.Evas_Object { return self.obj }

type Panel struct { obj *C.Evas_Object }
func (self *Panel) GetObj() *C.Evas_Object { return self.obj }

type Panes struct { obj *C.Evas_Object }
func (self *Panes) GetObj() *C.Evas_Object { return self.obj }

type Photocam struct { obj *C.Evas_Object }
func (self *Photocam) GetObj() *C.Evas_Object { return self.obj }

type Photo struct { obj *C.Evas_Object }
func (self *Photo) GetObj() *C.Evas_Object { return self.obj }

type Plug struct { obj *C.Evas_Object }
func (self *Plug) GetObj() *C.Evas_Object { return self.obj }

type Progressbar struct { obj *C.Evas_Object }
func (self *Progressbar) GetObj() *C.Evas_Object { return self.obj }

type Radio struct { obj *C.Evas_Object }
func (self *Radio) GetObj() *C.Evas_Object { return self.obj }

type Route struct { obj *C.Evas_Object }
func (self *Route) GetObj() *C.Evas_Object { return self.obj }

type SegmentControl struct { obj *C.Evas_Object }
func (self *SegmentControl) GetObj() *C.Evas_Object { return self.obj }

type Separator struct { obj *C.Evas_Object }
func (self *Separator) GetObj() *C.Evas_Object { return self.obj }

type Slider struct { obj *C.Evas_Object }
func (self *Slider) GetObj() *C.Evas_Object { return self.obj }

type Slideshow struct { obj *C.Evas_Object }
func (self *Slideshow) GetObj() *C.Evas_Object { return self.obj }

type Spinner struct { obj *C.Evas_Object }
func (self *Spinner) GetObj() *C.Evas_Object { return self.obj }

type Table struct { obj *C.Evas_Object }
func (self *Table) GetObj() *C.Evas_Object { return self.obj }

type Thumb struct { obj *C.Evas_Object }
func (self *Thumb) GetObj() *C.Evas_Object { return self.obj }

type Toolbar struct { obj *C.Evas_Object }
func (self *Toolbar) GetObj() *C.Evas_Object { return self.obj }

type Player struct { obj *C.Evas_Object }
func (self *Player) GetObj() *C.Evas_Object { return self.obj }

type Video struct { obj *C.Evas_Object }
func (self *Video) GetObj() *C.Evas_Object { return self.obj }

type Web struct { obj *C.Evas_Object }
func (self *Web) GetObj() *C.Evas_Object { return self.obj }

type Win struct { obj *C.Evas_Object }
func (self *Win) GetObj() *C.Evas_Object { return self.obj }

type ScrolledEntry struct { obj *C.Evas_Object }
func (self *ScrolledEntry) GetObj() *C.Evas_Object { return self.obj }
