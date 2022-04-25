package staticplugin

type Plugin interface {
	Exec()
}

var plugins = map[string]Plugin{}

func Register(name string, p Plugin) {
	plugins[name] = p
}

func Plugins() []Plugin {
	var r []Plugin
	for _, p := range plugins {
		r = append(r, p)
	}
	return r
}
