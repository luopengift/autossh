package cmd

import (
	"flag"
	"os"
	"strings"

	"github.com/luopengift/golibs/file"
	"github.com/luopengift/golibs/sys"
)

// Params params
type Params struct {
	Version  bool
	User     string
	Password string
	Port     string
	Key      string
	Hosts    []string
	Timeout  int
	Debug    bool
	Module   string
	Fork     int
	Args     string
}

// NewParams new params
func NewParams() *Params {
	version := flag.Bool("v", false, "(version)版本")
	user := flag.String("u", sys.Username(), "(username)用户名")
	password := flag.String("p", "", "(password)密码")
	port := flag.String("P", "22", "(Port)端口")
	key := flag.String("k", "", "(key)证书文件,绝对路径")
	iplist := flag.String("i", "", "IP地址列表,使用\",\"分割")
	ipfile := flag.String("file", "", "IP列表文件,使用\\n分割")
	timeout := flag.Int("t", 120, "(timeout)超时时间(单位:秒)")
	debug := flag.Bool("debug", false, "(debug)HTTP调试模式[http://debug(IP:PORT)/debug/pporf/]")
	module := flag.String("m", "", "(module)执行模块")
	args := flag.String("a", "", "(module_args)模块参数")
	fork := flag.Int("f", 5, "(fork)并发执行数")
	flag.Parse()

	var hosts []string
	if *ipfile != "" {
		st, _ := file.NewFile(*ipfile, os.O_RDONLY).ReadAll()
		hosts = strings.Split(strings.Trim(string(st), "\n"), "\n")
	} else {
		hosts = strings.Split(*iplist, ",")
	}

	return &Params{
		Version:  *version,
		User:     *user,
		Password: *password,
		Port:     *port,
		Key:      *key,
		Hosts:    hosts,
		Timeout:  *timeout,
		Debug:    *debug,
		Module:   *module,
		Fork:     *fork,
		Args:     *args,
	}
}
