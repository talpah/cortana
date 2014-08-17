package plugins

import "regexp"

type (
	Callback func(command string) (string, error)

	PluginManager interface {
		Register(command string, callback Callback, aliases map[string]*regexp.Regexp)
	}
)
