package hello

import (
	"fmt"
	"github.com/dlsniper/cortana/plugins"
	"regexp"
)

var (
	command = `Hello`
	aliases map[string]*regexp.Regexp
)

func isAlias(cmd string) bool {
	for _, alias := range aliases {
		if alias.MatchString(cmd) {
			return true
		}
	}

	return false
}

func canHandle(cmd string) bool {
	if cmd == command {
		return true
	}

	return isAlias(cmd)
}

func Initialize(pins plugins.PluginManager) {
	aliases = map[string]*regexp.Regexp{
		"Hi": regexp.MustCompile(`Hi`),
	}

	pins.Register(
		command,
		HelloWorld,
		aliases,
	)
}

func HelloWorld(cmd string) (string, error) {
	if !canHandle(cmd) {
		return "", fmt.Errorf("Can't handle command %s", cmd)
	}

	if isAlias(cmd) {
		return fmt.Sprintf("%s World!", cmd), nil
	}

	return "Hello World!", nil
}
