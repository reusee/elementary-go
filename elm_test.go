package elm

import (
  "testing"
  "fmt"
)

func TestBasicWin(t *testing.T) {
  win := NewWin(nil, "hello", WIN_BASIC)
  bg := NewBg(win)
  bg.SizeHintWeightSet(0.0, 0.0)
  bg.Show()

  win.Show()
  fmt.Printf("running\n")
  Run()
}
