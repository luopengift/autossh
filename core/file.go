package core

import (
	"os"
	"strings"
)

func read() {
	bu := make([]byte, 10000)
	f, _ := os.Open("hosts")

	n, _ := f.Read(bu)
	file := string(bu[:n])
	iplist := strings.Split(strings.Trim(file, "\n"), "\n")
	for _, ip := range iplist {
		//serverList.Add("", "", ip, 36000, "ops", "", "~/.ssh/xhqb.rsa")
		println(ip)
	}
}
