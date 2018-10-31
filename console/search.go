package console

import (
	"strings"

	"github.com/chzyer/readline"
	"github.com/luopengift/autossh/pkg/endpoint"
	"github.com/luopengift/log"
)

// 第三个参数表示是否要求返回唯一结果集.
func searchEndpoints(ins *readline.Instance, endpoints endpoint.Endpoints, one bool) (endpoint.Endpoints, error) {
	endpoints.Print()
	if len(endpoints) == 1 {
		return endpoints, nil
	}
	ins.SetPrompt(readline.StaticPrompt("ID/IP/主机> "))
	log.ConsoleWithGreen(`输入"ID/IP/主机"查询, 或者直接输入"s+ID/IP/主机"确认.`)
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	if isExit(input) {
		return nil, log.Errorf("exit")
	}
	if isNull(input) && !one {
		return nil, nil
	}
	var result endpoint.Endpoints
	if strings.HasPrefix(input, "s ") {
		result = endpoints.Match(strings.TrimPrefix(input, "s "))
	} else {
		inputList := strings.Split(input, " ")
		result = endpoints.Search(inputList...)
	}
	if !one {
		return result, nil
	}
	switch len(result) {
	case 0:
		return searchEndpoints(ins, endpoints, one)
	case 1:
		return result, nil
	default:
		return searchEndpoints(ins, result, one)
	}
}

func searchGroups(ins *readline.Instance, groups endpoint.Groups, one bool) (endpoint.Groups, error) {
	groups.Print()
	if len(groups) == 1 {
		return groups, nil
	}
	ins.SetPrompt(readline.StaticPrompt("ID/主机组> "))
	log.ConsoleWithGreen(`输入"ID/主机组"查询, 或者直接输入"s+ID/主机组"确认.`)
	input, err := ins.Readline()
	if err != nil {
		return nil, err
	}
	input = strings.TrimSpace(input)
	if isExit(input) {
		return nil, log.Errorf("exit")
	}
	if isNull(input) && !one {
		return nil, nil
	}

	var result endpoint.Groups
	if strings.HasPrefix(input, "s ") {
		result = groups.Match(strings.TrimPrefix(input, "s "))
	} else {
		inputList := strings.Split(input, " ")
		result = groups.Search(inputList...)
	}

	if !one {
		return result, nil
	}

	switch len(result) {
	case 0:
		return searchGroups(ins, groups, one)
	case 1:
		return result, nil
	default:
		return searchGroups(ins, result, one)
	}
}

func searchUsers(ins *readline.Instance, users endpoint.Users) ([]string, error) {
	if len(users) == 1 {
		return users, nil
	}
	users.Print()
	ins.SetPrompt(readline.StaticPrompt("ID/用户名> "))
	log.ConsoleWithGreen(`输入"ID/用户名"查询, 或者直接输入"s+ID/用户名"确认.`)
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
		result = users.Match(strings.TrimPrefix(input, "s "))
	} else {
		inputList := strings.Split(input, " ")
		result = users.Search(inputList...)
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

func searchExt(ins *readline.Instance, endpoints endpoint.Endpoints, groups endpoint.Groups) (endpoint.Endpoints, error) {
	var result endpoint.Endpoints

	for {
		ins.SetPrompt(readline.StaticPrompt("Batch>"))
		log.ConsoleWithGreen(`输入"P/g"选择主机列表, 按任意键确认.`)
		input, err := ins.Readline()
		if err != nil {
			return nil, err
		}
		input = strings.TrimSpace(input)
		switch {
		case input == "P", input == "p":
			endpointList, err := searchEndpoints(ins, endpoints, false)
			if err != nil {
				return nil, err
			}
			result = append(result, endpointList...)
		case input == "G", input == "g":
			groupList, err := searchGroups(ins, groups, false)
			if err != nil {
				return nil, err
			}
			for _, group := range groupList {
				result = append(result, group.Endpoints...)
			}
		default:
			log.ConsoleWithYellow("已选择的主机列表!")
			result.Print()
			return result, nil
		}
	}
}
