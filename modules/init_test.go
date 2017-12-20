package modules

import (
	"fmt"
	"testing"
)

func Test_paesr(t *testing.T) {
	str := "src=/ dest=/tmp"
	fmt.Println(str)
	parse(str)
	str = "src = / dest = /tmp"
	fmt.Println(str)
	parse(str)

	str = "src = / tt dest = /tmp"
	fmt.Println(str)
	parse(str)

}
