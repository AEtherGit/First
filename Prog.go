package main

import (
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
)

func main() {
  resp, _ := http.Get("https://mvideo.ru")
  bytes, _ := ioutil.ReadAll(resp.Body)
  string_body := string(bytes)
  mass := strings.Fields(string_body)
  for i, v := range mass {
    if v == "<a" {
      fmt.Println(mass[i])
    }
  }
//  fmt.Println(string_body)
  resp.Body.Close()
}
