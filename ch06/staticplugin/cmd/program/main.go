package main

import (
	"staticplugin"

	// プラグインをロード
	_ "staticplugin/plugins/a"
	_ "staticplugin/plugins/b"
)

func main() {
	for _, p := range staticplugin.Plugins() {
		p.Exec()
	}
}
