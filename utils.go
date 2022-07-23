package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lonzzi/BiliUpDynamicBot/e"
	"github.com/spf13/viper"
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

func Unicode2Str(raw string, threshold float64) (string, error) {
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

func IsExistDynamic(dynamics []BriefDynamic, dynamic BriefDynamic) bool {
	for _, v := range dynamics {
		if v.IDStr == dynamic.IDStr {
			return true
		}
	}
	return false
}

func AddNewDynamic(oldDynamics []BriefDynamic, dynamic BriefDynamic) ([]BriefDynamic, error) {
	if IsExistDynamic(oldDynamics, dynamic) {
		return oldDynamics, e.ERR_DYNAMIC_EXIST
	}

	oldDynamics = append(oldDynamics, dynamic)

	// 历史动态只保留10个
	if len(oldDynamics) > 10 {
		oldDynamics = oldDynamics[len(oldDynamics)-10:]
	}

	// 写入新数据
	fileContent, err := json.Marshal(oldDynamics)
	if err != nil {
		return nil, e.ERR_MARSHAL
	}
	if err := ioutil.WriteFile(viper.GetString("HistroyDynamicPath"), fileContent, 0644); err != nil {
		return nil, e.ERR_WRITEFILE
	}

	return oldDynamics, nil
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

func GetHistoryDynamics(filePath string) ([]BriefDynamic, error) {
	// 读出历史数据
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果文件不存在，则创建文件
			err = ioutil.WriteFile(filePath, []byte("[]"), 0666)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, e.ERR_READFILE
	}
	var oldDynamics []BriefDynamic
	if err := json.Unmarshal(fileContent, &oldDynamics); err != nil && len(fileContent) != 0 {
		return nil, e.ERR_UNMARSHAL
	}
	return oldDynamics, nil
}
