package utils

import (
	"net/http"
	"strings"
)

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
	for _, v := range strings.Split(s, "; ") {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			m[kv[0]] = kv[1]
		}
	}
	return m
}

func StrToCookies(s string) []*http.Cookie {
	var c []*http.Cookie
	for _, v := range strings.Split(s, "; ") {
		kv := strings.Split(v, "=")
		if len(kv) == 2 {
			c = append(c, &http.Cookie{Name: kv[0], Value: kv[1]})
		}
	}
	return c
}
