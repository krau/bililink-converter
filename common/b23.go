package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var client = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

func ConvertB23(b23code string) (string, error) {
	b23code = "https://b23.tv/" + b23code
	resp, err := client.Get(b23code)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 302 {
		var data map[string]any
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("invalid b23 link (status code: %d, message: %s", resp.StatusCode, data["message"])
	}
	convertedLink, err := resp.Location()
	if err != nil {
		return "", err
	}
	convertedLinkString := convertedLink.String()
	// 去除跟踪参数
	if strings.Contains(convertedLinkString, "?") {
		convertedLinkString = strings.Split(convertedLinkString, "?")[0]
	}
	return convertedLinkString, nil
}

func B23toAv(b23code string) (string, error) {
	convertedLink, err := ConvertB23(b23code)
	if err != nil {
		return "", err
	}
	bv := strings.TrimPrefix(convertedLink, "https://www.bilibili.com/video/")
	bv = strings.TrimSuffix(bv, "/")
	av, err := DecodeBV(bv)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://bilibili.com/video/av%d", av), nil
}

func B23toIframe(b23code string) (string, error) {
	convertedLink, err := ConvertB23(b23code)
	if err != nil {
		return "", err
	}
	bv := strings.TrimPrefix(convertedLink, "https://www.bilibili.com/video/")
	bv = strings.TrimSuffix(bv, "/")
	av, err := DecodeBV(bv)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://player.bilibili.com/player.html?aid=%d&autoplay=0", av), nil
}
