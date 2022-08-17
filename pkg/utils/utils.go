package utils

import (
	"io"
	"net/http"
	"strings"

	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/spf13/viper"
)

func Fetch(url string, cookie ...*http.Cookie) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", viper.GetString("server.user_agent"))
	if len(cookie) > 0 {
		req.AddCookie(cookie[0])
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, e.ERR_HTTP_STATUS_NOT_OK
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CookieToString(c []*http.Cookie) string {
	var s []string
	for _, v := range c {
		s = append(s, v.Name+"="+v.Value)
	}
	return strings.Join(s, "; ")
}

func CookieToMap(c []*http.Cookie) map[string]string {
	m := make(map[string]string)
	for _, v := range c {
		m[v.Name] = v.Value
	}
	return m
}

func StrToMap(s string) map[string]string {
	m := make(map[string]string)
	s = strings.ReplaceAll(s, " ", "")
	for _, v := range strings.Split(s, ";") {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			m[kv[0]] = kv[1]
		}
	}
	return m
}
