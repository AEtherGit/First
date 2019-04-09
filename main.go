package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://mvideo.ru")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	stringBody := string(bytes)
	mass := strings.Fields(stringBody)
	for i, v := range mass {
		if v == "<a" {
			fmt.Println(mass[i])
		}
	}
	//  fmt.Println(string_body)
	resp.Body.Close()
}
