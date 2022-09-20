package plugin

import (
	"log"
	"plugin"
)

type PluginFunc func(string) (string, error)

var pluginPath map[string]string = map[string]string{
	"UnicodeToStr": "plugin/unicodeToStr.so",
}

var localPlugin map[string]PluginFunc = make(map[string]PluginFunc)

func LoadLocalPlugins() {
	for name, path := range pluginPath {
		plug, err := plugin.Open(path)
		if err != nil {
			panic(err)
		}

		middleFunc, err := plug.Lookup("ExecuteFunc")
		if err != nil {
			panic(err)
		}

		GenFunc, ok := middleFunc.(func() PluginFunc)
		if !ok {
			log.Panic("This plug", plug, "has some errors")
		}

		localPlugin[name] = GenFunc()
	}
}

func GetFunc(name string) PluginFunc {
	return localPlugin[name]
}
