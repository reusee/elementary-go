package elm

import (
  "testing"
)

func TestBasicWin(t *testing.T) {
  win := NewWin(nil, "hello", WIN_BASIC)
  win.Show()
}
