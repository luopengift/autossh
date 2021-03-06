package console

func include(input string, items ...string) bool {
	for _, v := range items {
		if input == v {
			return true
		}
	}
	return false
}

func isNull(str string) bool {
	return str == ""
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
