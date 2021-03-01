package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(foo([]string{"a", "b"}))
}

func foo(tags []string) string {
selections := make([]string, len(tags))

	for i := range selections {
		clause := "SERVER.ID = SERVER_TAG.SERVERID AND SELECT SERVER.ID FROM SERVER, SERVER_TAG WHERE SERVER_TAG.TAGID = ? SERVER.ID = SERVER_TAG.SERVERID"
		selections[i] = clause
	}

	return strings.Join(selections, " intersect ")
}
