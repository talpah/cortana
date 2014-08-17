package pluginmanager

import (
	"fmt"
	"github.com/dlsniper/cortana/plugins"
	"github.com/dlsniper/cortana/plugins/echo"
	"github.com/dlsniper/cortana/plugins/hello"
	"regexp"
)

type (
	PluginManager struct {
		canonicalCommands map[string]plugins.Callback
		aliases           map[*regexp.Regexp]plugins.Callback
		commands          map[*regexp.Regexp]plugins.Callback
	}
)

func (pm *PluginManager) findCommand(command string) (plugins.Callback, error) {
	for cmd, callback := range pm.commands {
		if cmd.MatchString(command) {
			return callback, nil
		}
	}

	for cmd, callback := range pm.aliases {
		if cmd.MatchString(command) {
			return callback, nil
		}
	}

	return nil, fmt.Errorf("command '%s' not found", command)
}

func (pm *PluginManager) Register(canonicalCommand, command string, callback, help plugins.Callback, aliases map[string]*regexp.Regexp) {
	pm.canonicalCommands[canonicalCommand] = help

	regex := regexp.MustCompile(command)

	if _, ok := pm.commands[regex]; ok {
		panic(fmt.Errorf("command '%s' is already registered", command))
	}

	pm.commands[regex] = callback

	for alias, compiledAlias := range aliases {
		if _, ok := pm.aliases[compiledAlias]; ok {
			panic(fmt.Errorf("alias '%s' is already registered", alias))
		}

		pm.aliases[compiledAlias] = callback
	}
}

func (pm *PluginManager) Execute(command string) (string, error) {
	cmd, err := pm.findCommand(command)
	if err != nil {
		return "", fmt.Errorf("command '%s' not found", command)
	}

	return cmd(command)
}

func (pm *PluginManager) Help(command string) (string, error) {
	regex := regexp.MustCompile(`(?i)help(?:\s)?(.*)`)
	result := regex.FindStringSubmatch(command)

	if result[1] == "" {
		return "Generic help", nil
	}

	if _, ok := pm.canonicalCommands[result[1]]; ok {
		return pm.canonicalCommands[result[1]](command)
	}

	return fmt.Sprintf("Help for %s not found.", command), nil
}

func (pm *PluginManager) Initialize() {
	pm.canonicalCommands = make(map[string]plugins.Callback)
	pm.aliases = make(map[*regexp.Regexp]plugins.Callback)
	pm.commands = make(map[*regexp.Regexp]plugins.Callback)

	pm.Register(
		"help",
		`(?:(?i)help)(?:\s)?(.*)`,
		pm.Help,
		pm.Help,
		make(map[string]*regexp.Regexp),
	)

	hello.Initialize(pm)
	echo.Initialize(pm)
}
