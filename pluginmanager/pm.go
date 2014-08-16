package pluginmanager

import (
	"fmt"
	"github.com/dlsniper/cortana/plugins"
	"github.com/dlsniper/cortana/plugins/hello"
)

type (
	PluginManager struct {
		aliases map[string]plugins.Callback
		commands map[string]plugins.Callback
	}
)

func (pm *PluginManager) findCommand(command string) (plugins.Callback, error) {
	_, ok := pm.commands[command]
	if ok {
		return pm.commands[command], nil
	}

	_, ok = pm.aliases[command]
	if ok {
		return pm.aliases[command], nil
	}

	return nil, fmt.Errorf("command '%s' not found", command)
}

func (pm *PluginManager) Register(command string, callback plugins.Callback, aliases map[int]string) {
	if _, ok := pm.commands[command]; ok {
		panic(fmt.Errorf("command '%s' is already registered", command))
	}

	pm.commands[command] = callback

	for _, alias := range aliases {
		if _, ok := pm.aliases[alias]; ok {
			panic(fmt.Errorf("alias '%s' is already registered for command '%s'", alias, pm.aliases[alias]))
		}
		pm.aliases[alias] = callback
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
	pm.aliases = make(map[string]plugins.Callback)
	pm.commands = make(map[string]plugins.Callback)

	hello.Initialize(pm)
}
