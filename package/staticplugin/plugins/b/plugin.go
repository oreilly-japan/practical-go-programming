package b

import (
	"fmt"

	"staticplugin"
)

type pluginB struct {
}

func (p pluginB) Exec() {
	fmt.Println("plugin B")
}

func init() {
	staticplugin.Register("B", &pluginB{})
}
