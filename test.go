package main

import (
	"fmt"
	"strings"
)

func main() {
	content := "42.118.242.143/26"
	fmt.Println (strings.Split(content, ". /"))
}
