package echo

import (
	"fmt"
	"github.com/dlsniper/cortana/plugins"
	"regexp"
)

var (
	command = `(?:E|e)cho(?:\s)?(.*)`
	compiledCmd = regexp.MustCompile(command)
	aliases map[string]*regexp.Regexp
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
	aliases = map[string]*regexp.Regexp{
	}

	pins.Register(
		command,
		Echo,
		aliases,
	)
}

func Echo(cmd string) (string, error) {
	regex := canHandle(cmd)
	if regex == nil {
		return "", fmt.Errorf("Can't handle command %s", cmd)
	}

	result := regex.FindStringSubmatch(cmd)

	return result[1], nil
}
