package console

import (
	"regexp"
	"strings"

	"github.com/luopengift/ssh"
)

func include(input string, items ...string) bool {
	for _, v := range items {
		if input == v {
			return true
		}
	}
	return false
}

func isExit(str string) bool {
	return include(str, "q", "Q", "quit", "exit")
}

func isVersion(str string) bool {
	return include(str, "V", "v", "version")
}

func isHelp(str string) bool {
	return include(str, "h", "H", "help")
}

func getEndpoint(str string) *ssh.Endpoint {
	endpoint := ssh.NewEndpoint()
	if strings.Contains(str, "@") {
		d := strings.Split(str, "@")
		endpoint.User, endpoint.IP = d[0], d[1]
	} else {
		endpoint.IP = str
	}
	if strings.HasPrefix(endpoint.IP, "ip-") { // support aws hostname format
		re := regexp.MustCompile(`[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}`)
		endpoint.IP = strings.Replace(re.FindString(endpoint.IP), "-", ".", -1)
	}
	return endpoint
}
