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
	ipList   string
	ipFiles  string
	Group    string
	Timeout  int
	Pseudo   bool
	Debug    bool
	Batch    bool
	Module   string
	Fork     int
	Labels   string
	Args     string
}

// Hosts get hosts
func (params *Params) Hosts() ([]string, error) {
	var hosts []string

	if params.ipList != "" {
		hosts = strings.Split(params.ipList, ",")
	}

	if params.ipFiles != "" {
		for _, ipFile := range strings.Split(params.ipFiles, ";") {
			st, err := file.NewFile(ipFile, os.O_RDONLY).ReadAll()
			if err != nil {
				return nil, err
			}
			hosts = strings.Split(strings.TrimSpace(string(st)), "\n")
		}
	}
	return hosts, nil
}

// NewParams new params
func NewParams() *Params {
	version := flag.Bool("v", false, "(version)版本")
	user := flag.String("u", sys.Username(), "(username)用户名")
	password := flag.String("p", "", "(password)密码")
	port := flag.String("P", "22", "(Port)端口")
	key := flag.String("k", "", "(key)证书文件,绝对路径")
	ipList := flag.String("i", "", `IP地址列表,使用","分割`)
	ipFiles := flag.String("files", "", `IP列表文件,使用"\n"分格,多个文件用";"区分`)
	group := flag.String("g", "", "(group)主机组")
	timeout := flag.Int("timeout", 120, "(timeout)超时时间(单位:秒)")
	pseudo := flag.Bool("t", false, "(pseudo)pseudo-terminal allocation")
	debug := flag.Bool("debug", false, "(debug)HTTP调试模式[http://debug(IP:PORT)/debug/pporf/]")
	batch := flag.Bool("b", false, "(batch)批量执行模式")
	module := flag.String("m", "", "(module)执行模块")
	args := flag.String("a", "", "(module_args)模块参数")
	labels := flag.String("l", "", "(labels)索引标签")
	fork := flag.Int("f", 5, "(fork)并发执行数")
	flag.Parse()

	return &Params{
		Version:  *version,
		User:     *user,
		Password: *password,
		Port:     *port,
		Key:      *key,
		ipList:   *ipList,
		ipFiles:  *ipFiles,
		Group:    *group,
		Timeout:  *timeout,
		Pseudo:   *pseudo,
		Debug:    *debug,
		Batch:    *batch,
		Module:   *module,
		Fork:     *fork,
		Labels:   *labels,
		Args:     *args,
	}
}
