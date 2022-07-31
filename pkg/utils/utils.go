package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/lonzzi/bilibot/pkg/e"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Fetch Error: ", err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("Error: ", resp.StatusCode)
		return nil, e.ERR_HTTP_STATUS_NOT_OK
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}

func UnicodeToStr(raw string, threshold float64) (string, error) {
	if threshold < 0 || threshold > 1 {
		return raw, e.ERR_INVALID_NUMBER
	}
	quoteStr := strconv.Quote(raw)
	log.Println("含Unicode浓度: ", float64(strings.Count(quoteStr, `\\u`))/float64(len(quoteStr)))
	if float64(strings.Count(quoteStr, `\\u`))/float64(len(quoteStr)) < threshold {
		return raw, e.ERR_BELOW_THRESHOLD
	}
	str, err := strconv.Unquote(strings.Replace(quoteStr, `\\u`, `\u`, -1))
	if err != nil {
		log.Fatal(err)
	}
	return str, nil
}

func StrToUnicode(str string) string {
	DD := []rune(str)
	finallStr := ""
	for i := 0; i < len(DD); i++ {
		if unicode.Is(unicode.Scripts["Han"], DD[i]) {
			textQuoted := strconv.QuoteToASCII(string(DD[i]))
			finallStr += textQuoted[1 : len(textQuoted)-1]
		} else {
			h := fmt.Sprintf("%x", DD[i])
			finallStr += "\\u" + isFullFour(h)
		}
	}
	return finallStr
}

func isFullFour(str string) string {
	if len(str) == 1 {
		str = "000" + str
	} else if len(str) == 2 {
		str = "00" + str
	} else if len(str) == 3 {
		str = "0" + str
	}
	return str
}

func StrUrlToMap(params []string) map[string]string {
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
