package main

import (
	"fmt"
	"github.com/wagoodman/keybinding"
)

func show(keyStr string) {
	keys, err := keybinding.ParseAll(keyStr)
	if err != nil {
		fmt.Println("Error parsing", keyStr, ":", err)
	} else {
		fmt.Println("Key: ", keyStr, "=", keys)
	}
}

func main() {
	show("ctrl+a")
	show("ctrl+b")
	show("ctrl+/, tab")
	show("ctrl+   alt +/")
}
