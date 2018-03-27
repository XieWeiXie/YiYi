package utils

import (
	"YiYi/domain/httpclient"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
)

func StringReplacer(value string) string {
	newData := strings.NewReplacer(`<p class="fr-pspace-a0">`, "", "<p>", "", "</p>", "", "<br>", "",
		`<div class="one-img-container one-img-container-no-note">`, "",
		"</div>", "")
	newString := newData.Replace(value)
	return newString
}

func Response(url string, key string) gjson.Result {
	var newClient httpclient.HttpClient
	newClient = httpclient.Implement{}
	resp, err := newClient.Get(url)
	//fmt.Println(url)
	if err != nil {
		fmt.Println("Get HTTP Response Failed")
		return gjson.Result{}
	}
	if key == "" {
		return gjson.Parse(string(resp))
	} else {
		return gjson.Parse(string(resp)).Get(key)
	}
}

func ListConvertString(args []string) string {
	return strings.Join(args, " | ")
}
