package util

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func GetGrassSVG(username string) string {
	url := "https://github.com/users/" + username + "/contributions"
	resp, _ := http.Get(url)

	byteArray, _ := ioutil.ReadAll(resp.Body)
	exp := regexp.MustCompile(`(?s)<svg(.*)<\/svg>`)
	svg := exp.FindAllStringSubmatch(string(byteArray), -1)[0][0]

	// svgのスタイルを調整する
	svg = strings.Replace(svg, `class="month"`, `fill="#767676" font-size="9"`, -1)
	svg = strings.Replace(svg, `class="wday"`, `fill="#767676" font-size="9"`, -1)
	return svg
}
