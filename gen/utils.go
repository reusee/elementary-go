package main

import (
  "fmt"
)

func p(format string, args ...interface{}) {
  fmt.Printf(format, args...)
}
