package elm

import (
  "testing"
  "fmt"
)

func TestBasicWin(t *testing.T) {
  win := NewWin(nil, "hello", WIN_BASIC)
  win.TitleSet("Hello, world!")

  bg := NewBg(win)
  bg.SizeHintWeightSet(0.0, 0.0)
  win.ResizeObjectAdd(bg)
  bg.Show()

  box := NewBox(win)
  box.HorizontalSet(true)
  box.SizeHintWeightSet(0.0, 0.0)
  win.ResizeObjectAdd(box)
  box.Show()

  lab := NewLabel(win)
  lab.PartTextSet("", "FOO")
  lab.SizeHintWeightSet(0.0, 0.0)
  box.PackEnd(lab)
  lab.Show()

  btn := NewButton(win)
  btn.PartTextSet("", "OK")
  btn.SizeHintWeightSet(0.0, 0.0)
  box.PackEnd(btn)
  btn.Show()

  win.Show()
  fmt.Printf("running\n")
  Run()
}
