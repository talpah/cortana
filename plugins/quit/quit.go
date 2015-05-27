package quit

import (
	"fmt"
	"../../plugins"
	"regexp"
	"os"
)

var (
	canonicalCommand = "quit"
	command          = `(Q|q)uit`
	aliases          map[string]*regexp.Regexp
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
	if regexp.MustCompile(command).MatchString(cmd) {
		return true
	}

	return isAlias(cmd)
}

func Initialize(pins plugins.PluginManager) {
	aliases = map[string]*regexp.Regexp{
		"Exit": regexp.MustCompile(`(E|e)xit`),
//		"Quit": regexp.MustCompile(`quit`),
	}

	pins.Register(
		canonicalCommand,
		command,
		Quit,
		Help,
		aliases,
	)
}

func Help(cmd string) (string, error) {
	return `Quit command help

Quit will... quit

Usage example:
quit
Quit
exit
Exit
`, nil
}

func Quit(cmd string) (string, error) {
	if !canHandle(cmd) {
		return "", fmt.Errorf("Can't handle command %s", cmd)
	}
	os.Exit(0)
	return "", nil
}
