package plugins

import "regexp"

type (
	Callback func(command string) (string, error)

	PluginManager interface {
		Register(canonicalCommand, command string, callback, help Callback, aliases map[string]*regexp.Regexp)
	}
)
