package console

import (
	"strings"

	"github.com/chzyer/readline"
	"github.com/luopengift/autossh/pkg/endpoint"
	"github.com/luopengift/log"
)

func searchEndpoints(ins *readline.Instance, endpoints endpoint.Endpoints) (endpoint.Endpoints, error) {
	endpoints.Print()
	if len(endpoints) == 1 {
		return endpoints, nil
	}
	ins.SetPrompt(readline.StaticPrompt("ID/IP/主机> "))
	log.ConsoleWithGreen(`输入"s+ID/IP/主机"查询, 或者直接输入"ID/IP/主机"确认.`)
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	if isExit(input) {
		return nil, log.Errorf("exit")
	}

	var result endpoint.Endpoints
	if strings.HasPrefix(input, "s ") {
		inputList := strings.Split(strings.TrimPrefix(input, "s "), " ")
		result = endpoints.Search(inputList...)
	} else {
		result = endpoints.Match(input)
	}

	switch len(result) {
	case 0:
		return searchEndpoints(ins, endpoints)
	case 1:
		return result, nil
	default:
		return searchEndpoints(ins, result)
	}

}

func searchGroups(ins *readline.Instance, groups *endpoint.Groups) (*endpoint.Groups, error) {
	groups.Print()
	if len(groups.List) == 1 {
		return groups, nil
	}
	ins.SetPrompt(readline.StaticPrompt("ID/主机组> "))
	log.ConsoleWithGreen(`输入"s+ID主机组"查询, 或者直接输入"ID/主机组"确认.`)
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	if isExit(input) {
		return nil, log.Errorf("exit")
	}

	var result *endpoint.Groups
	if strings.HasPrefix(input, "s ") {
		inputList := strings.Split(strings.TrimPrefix(input, "s "), " ")
		result = groups.Search(inputList...)
	} else {
		result = groups.Match(input)
	}

	switch len(result.List) {
	case 0:
		return searchGroups(ins, groups)
	case 1:
		return result, nil
	default:
		return searchGroups(ins, result)
	}
}

func searchUsers(ins *readline.Instance, users endpoint.Users) ([]string, error) {
	if len(users) == 1 {
		return users, nil
	}
	users.Print()
	ins.SetPrompt(readline.StaticPrompt("ID/用户名> "))
	log.ConsoleWithGreen(`输入"s+ID/用户名"查询, 或者直接输入"ID/用户名"确认.`)
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	if isExit(input) {
		return nil, log.Errorf("exit")
	}

	var result endpoint.Users
	if strings.HasPrefix(input, "s ") {
		inputList := strings.Split(strings.TrimPrefix(input, "s "), " ")
		result = users.Search(inputList...)
	} else {
		result = users.Match(input)
	}

	switch len(result) {
	case 0:
		return searchUsers(ins, users)
	case 1:
		return result, nil
	default:
		return searchUsers(ins, result)
	}
}
