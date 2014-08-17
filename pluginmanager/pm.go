package pluginmanager

import (
	"fmt"
	"github.com/dlsniper/cortana/plugins"
	"github.com/dlsniper/cortana/plugins/hello"
	"regexp"
	"github.com/dlsniper/cortana/plugins/echo"
)

type (
	PluginManager struct {
		aliases  map[*regexp.Regexp]plugins.Callback
		commands map[*regexp.Regexp]plugins.Callback
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

func (pm *PluginManager) Register(command string, callback plugins.Callback, aliases map[string]*regexp.Regexp) {
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

func (pm *PluginManager) Initialize() {
	pm.aliases = make(map[*regexp.Regexp]plugins.Callback)
	pm.commands = make(map[*regexp.Regexp]plugins.Callback)

	hello.Initialize(pm)
	echo.Initialize(pm)
}
