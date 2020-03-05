package util

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func GetGrassSVG(username string) string {
	url := "https://github.com/users/" + username + "/contributions"
	resp, _ := http.Get(url)

	byteArray, _ := ioutil.ReadAll(resp.Body)
	exp := regexp.MustCompile(`(?s)<svg(.*)<\/svg>`)
	svg := exp.FindAllStringSubmatch(string(byteArray), -1)[0][0]
	return svg
}

func main() {
	fmt.Println(GetGrassSVG("mkan0141"))
}
