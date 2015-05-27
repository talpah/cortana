package echo

import (
	"fmt"
	"../../plugins"
	"regexp"
)

var (
	canonicalCommand = "echo"
	command          = `(?:E|e)cho(?:\s)?(.*)`
	compiledCmd      = regexp.MustCompile(command)
	aliases          map[string]*regexp.Regexp
)

func isAlias(cmd string) *regexp.Regexp {
	for _, alias := range aliases {
		if alias.MatchString(cmd) {
			return alias
		}
	}

	return nil
}

func canHandle(cmd string) *regexp.Regexp {
	if compiledCmd.MatchString(cmd) {
		return compiledCmd
	}

	return isAlias(cmd)
}

func Initialize(pins plugins.PluginManager) {
	aliases = map[string]*regexp.Regexp{}

	pins.Register(
		canonicalCommand,
		command,
		Echo,
		Help,
		aliases,
	)
}

func Help(cmd string) (string, error) {
	return `Echo command help

Echo is a more advanced command.
Echo will always reply back everything what's after 'echo', making this a good mirror service.

Usage example:
echo Demo
echo Hello World
`, nil
}

func Echo(cmd string) (string, error) {
	regex := canHandle(cmd)
	if regex == nil {
		return "", fmt.Errorf("Can't handle command %s", cmd)
	}

	result := regex.FindStringSubmatch(cmd)

	return result[1], nil
}
