// go build -buildmode=plugin -o plugin/unicodeToStr.so plugin/unicodeToStr.go
package main

import (
	"strconv"
	"strings"

	"github.com/Augenblick-tech/bilibot/pkg/plugin"
)

func UnicodeToStr(raw string) (string, error) {
	quoteStr := strconv.Quote(raw)
	if strings.Count(quoteStr, `\\u`) == 0 {
		return "", nil
	}
	str, err := strconv.Unquote(strings.Replace(quoteStr, `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

func ExecuteFunc() plugin.PluginFunc {
	return plugin.PluginFunc(UnicodeToStr)
}
