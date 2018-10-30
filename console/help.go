package console

import (
	"time"

	"github.com/luopengift/log"
)

func help() {
	log.ConsoleWithBlue("### 欢迎使用Autossh Jump System[%s] ###", time.Now().Format("2006/01/02 15:04:05"))
	log.ConsoleWithGreen("")
	for idx, v := range []string{
		"输入 P/p 查看主机列表.",
		"输入 G/g 查看主机分组",
		"输入 E/e 批量执行命令",
		"输入 s + IP 直接登录.",
		"输入 V/v 查看版本号.",
		"输入 H/h 帮助.",
		"输入 Q/q 退出.",
	} {
		log.ConsoleWithGreen("\t%d) %s", idx, v)
	}
	log.ConsoleWithGreen("")
}
