package plugins

type (
	Callback func(command string) (string, error)

	PluginManager interface {
		Register(command string, callback Callback, aliases map[int]string)
	}
)
