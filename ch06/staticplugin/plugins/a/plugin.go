package a

import (
	"fmt"

	"staticplugin"
)

type pluginA struct {
}

func (p pluginA) Exec() {
	fmt.Println("plugin A")
}

func init() {
	staticplugin.Register("a", &pluginA{})
}
