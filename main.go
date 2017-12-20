package main

import (
	"fmt"
	"github.com/luopengift/autossh/cmd"
)

func main() {
	if err := cmd.Exec(); err != nil {
		fmt.Println(err)
	}
}
