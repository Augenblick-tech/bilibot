package main

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/lonzzi/BiliUpDynamicBot/e"
)

func Unicode2Str(raw string, threshold float64) (string, error) {
	if threshold < 0 || threshold > 1 {
		return raw, errors.New(e.InvalidNumber)
	}
	quoteStr := strconv.Quote(raw)
	log.Println("threshold: ", float64(strings.Count(quoteStr, `\\u`))/float64(len(quoteStr)))
	if float64(strings.Count(quoteStr, `\\u`))/float64(len(quoteStr)) < threshold {
		return raw, errors.New(e.LessThreshold)
	}
	str, err := strconv.Unquote(strings.Replace(quoteStr, `\\u`, `\u`, -1))
	if err != nil {
		panic(err)
	}
	return str, nil
}

func IsExistDynamic(dynamics []BriefDynamic, dynamic BriefDynamic) bool {
	for _, v := range dynamics {
		if v.IDStr == dynamic.IDStr {
			return true
		}
	}
	return false
}

func AddNewDynamic(dynamics []BriefDynamic, dynamic BriefDynamic) ([]BriefDynamic, error) {
	if IsExistDynamic(dynamics, dynamic) {
		return dynamics, errors.New(e.DynamicAlreadyExist)
	}

	dynamics = append(dynamics, dynamic)
	if len(dynamics) == 10 {
		dynamics = dynamics[1:]
	}

	return dynamics, nil
}

func StrUrl2Map(params []string) map[string]string {
	m := make(map[string]string)
	for _, v := range params {
		kv := strings.Split(v, "=")
		if len(kv) != 2 {
			continue
		}
		m[kv[0]] = kv[1]
	}
	return m
}