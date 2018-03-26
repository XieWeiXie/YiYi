package one

import (
	"demo_for_golang/demo-for-cli/domain/httpclient"
	"demo_for_golang/demo-for-cli/utils"

	"fmt"

	"demo_for_golang/demo-for-cli/objects"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

const (
	ListAPI    = "http://v3.wufazhuce.com:8000/api/onelist/idlist"
	ListOneAPI = "http://v3.wufazhuce.com:8000/api/onelist/%s/0"
	DetailAPI  = "http://v3.wufazhuce.com:8000/api/essay/%s"
)

var StoryUsage = "get story info from One API"

func SubStoryCommand() []cli.Command {
	return []cli.Command{
		{
			Name:   "IdList",
			Action: actionIDList,
		},
		{
			Name:        "OneList",
			Subcommands: subCommandOneList(),
		},
		{
			Name:        "Detail",
			Subcommands: subCommandOneDetail(),
		},
	}
}

func subCommandOneList() []cli.Command {
	return []cli.Command{
		{
			Name:   "weather",
			Usage:  "get weather",
			Action: actionGetOneListWeather,
		},
		{
			Name:   "menu",
			Usage:  "get menu",
			Action: actionGetOneListMenu,
		},
	}

}

func subCommandOneDetail() []cli.Command {
	return []cli.Command{
		{
			Name:   "info",
			Action: actionOneDetailInfo,
		},
		{
			Name:   "content",
			Action: actionOneDetailContent,
		},
		{
			Name:   "share",
			Action: actionOneDetailShare,
		},
	}

}

func actionIDList(c *cli.Context) {
	if c.NArg() != 0 {
		fmt.Println("You need not add one argument")
		return
	}
	if c.NArg() == 0 {
		getIDList(ListAPI)
	}
}
func actionGetOneListWeather(c *cli.Context) {
	if c.NArg() == 1 {
		getOneListWeather(fmt.Sprintf(ListOneAPI, c.Args().Get(0)))
	}
}
func actionGetOneListMenu(c *cli.Context) {
	if c.NArg() == 1 {
		getOneListMenu(fmt.Sprintf(ListOneAPI, c.Args().Get(0)))
	}
}

func actionOneDetailInfo(c *cli.Context) {
	if c.NArg() == 1 {
		getOneDetailInfo(fmt.Sprintf(DetailAPI, c.Args().Get(0)))
	}
}

func actionOneDetailContent(c *cli.Context) {
	if c.NArg() == 1 {
		getDetailContent(fmt.Sprintf(DetailAPI, c.Args().Get(0)))
	}

}

func actionOneDetailShare(c *cli.Context) {
	if c.NArg() == 1 {
		getDetailShare(fmt.Sprintf(DetailAPI, c.Args().Get(0)))
	}
}

// 获取所有id-list
func getIDList(url string) {
	var newClient httpclient.HttpClient
	newClient = httpclient.Implement{}
	resp, err := newClient.Get(url)
	if err != nil {
		fmt.Println("Get HTTP Response Failed")
		return
	}
	data := gjson.Parse(string(resp))
	rs := objects.ListData{
		Response: data.Get("res").String(),
		Data:     data.Get("data").Raw,
	}
	fmt.Println(rs)
}

// 获取当天天气情况
func getOneListWeather(url string) {
	data := commonURLResponse(url, "data.weather")
	weather := objects.Weather{
		CityName:      data.Get("city_name").String(),
		Date:          data.Get("date").String(),
		Temperature:   data.Get("temperature").String(),
		Humidity:      data.Get("humidity").String(),
		Climate:       data.Get("climate").String(),
		WindDirection: data.Get("wind_direction").String(),
		Hurricane:     data.Get("hurricane").String(),
	}
	fmt.Println(weather)

}

// 获取当天的七个栏目简介
func getOneListMenu(url string) {
	data := commonURLResponse(url, "data.menu")
	menu := objects.Menu{}
	allListInfo := []objects.ListInfo{}
	menu.Vol = data.Get("vol").String()
	for _, one := range data.Get("list").Array() {
		oneMap := make(map[string]string)
		oneMap["id"] = one.Get("tag.id").String()
		oneMap["title"] = one.Get("tag.title").String()
		allListInfo = append(allListInfo, objects.ListInfo{
			ContentType: one.Get("content_type").String(),
			ContentID:   one.Get("content_id").String(),
			Title:       one.Get("title").String(),
			Tag:         oneMap,
		})
	}
	menu.List = allListInfo
	menu.Template()
}

// 获取某个栏目的信息
func getOneDetailInfo(url string) {
	data := commonURLResponse(url, "data")
	contentInfo := objects.ContentInfo{}
	contentInfo.ID = data.Get("content_id").String()
	contentInfo.Title = data.Get("hp_title").String()
	contentInfo.Author = objects.Author{
		UserID:    data.Get("author").Array()[0].Get("user_id").String(),
		UserName:  data.Get("author").Array()[0].Get("user_name").String(),
		Desc:      data.Get("author").Array()[0].Get("desc").String(),
		WbName:    data.Get("author").Array()[0].Get("wb_name").String(),
		Summary:   data.Get("author").Array()[0].Get("summary").String(),
		FansTotal: data.Get("author").Array()[0].Get("fans_total").String(),
	}
	contentInfo.Auth = data.Get("auth_it").String()
	contentInfo.AuthorIntroduce = data.Get("hp_author_introduce").String()
	contentInfo.MakeTime = data.Get("hp_makettime").String()
	contentInfo.LastUpdateDate = data.Get("last_update_date").String()
	contentInfo.WebURL = data.Get("web_url").String()
	contentInfo.GuideWord = data.Get("guide_word").String()
	contentInfo.Audio = data.Get("audio").String()
	contentInfo.Anchor = data.Get("anchor").String()
	contentInfo.EditorEmail = data.Get("editor_email").String()
	contentInfo.Template()
}

// 获取某个栏目的具体内容
func getDetailContent(url string) {
	data := commonURLResponse(url, "data")
	fmt.Println(utils.StringReplacer(data.Get("hp_content").String()))

}

// 获取某个栏目的分享信息
func getDetailShare(url string) {
	data := commonURLResponse(url, "data.share_list")
	var shareInfo objects.ShareInfo
	var keyList = []string{"wx", "wx_timeline", "weibo", "qq"}
	for _, one := range keyList {
		shareInfo.ShareList = append(shareInfo.ShareList, handleShareInfo(one, data))
	}
	fmt.Println(shareInfo)

}

func commonURLResponse(url string, keyValue string) gjson.Result {
	var newClient httpclient.HttpClient
	newClient = httpclient.Implement{}
	resp, err := newClient.Get(url)
	if err != nil {
		fmt.Println("Get HTTP Response Failed")
		return gjson.Result{}
	}
	data := gjson.Parse(string(resp)).Get(keyValue)
	return data
}

func handleShareInfo(key string, args gjson.Result) map[string]objects.ShareDetail {
	var value map[string]objects.ShareDetail
	value = make(map[string]objects.ShareDetail)
	value[key] = objects.ShareDetail{
		Title:    args.Get(key).Get("title").String(),
		Desc:     args.Get(key).Get("desc").String(),
		Link:     args.Get(key).Get("link").String(),
		ImageURL: args.Get(key).Get("imgUrl").String(),
	}
	return value
}
