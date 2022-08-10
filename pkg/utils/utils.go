package utils

import (
	"io"
	"net/http"

	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/spf13/viper"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", viper.GetString("server.user_agent"))
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
