package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "This is a big apple tree. I love big big apple! 42"
    ss := strings.Fields(text)
    for i, v := range ss {
      if v == "big" {
      fmt.Println(ss[i])
    }
    }
    m := ss[3]
    fmt.Println(m)
}
