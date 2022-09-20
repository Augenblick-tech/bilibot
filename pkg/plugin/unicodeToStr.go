package plugin

import (
	"strconv"
	"strings"
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
