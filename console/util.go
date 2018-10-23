package console

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
