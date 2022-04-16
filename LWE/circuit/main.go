package main

import (
  "fmt"
)

func main() {
  for i := 0; i < 0b100; i++ {
    for j := 0; j < 0b100; j++ {
      if (i*j)&2 != 2 {
        fmt.Printf("%.2b %.2b %.4b\n", i,j,i*j)

      }
    }
  }
}
