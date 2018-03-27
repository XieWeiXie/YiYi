package douban

import (
	"YiYi/utils"
	"fmt"

	"YiYi/objects"

	"math/rand"
	"strconv"

	"time"

	"strings"

	"github.com/urfave/cli"
)

const (
	bookApi           = "https://api.douban.com/v2/book/series/%s/books"
	bookDetail        = "https://api.douban.com/v2/book/isbn/%s"
	bookSearchByQuery = "https://api.douban.com/v2/book/search?q=%s&count=%s"
	bookSearchByTag   = "https://api.douban.com/v2/book/search?tag=%s&count=%s"
)

var BookUsage = "get book info from DouBan API"

func getFlagSearch() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:  "count, c",
			Value: 20,
			Usage: "get count for search response",
		},
	}
}
func SubBookCommand() []cli.Command {
	return []cli.Command{
		{
			Name:   "random",
			Action: actionBookNumber,
		},
		{
			Name:   "detail",
			Action: actionBookDetail,
		},
		{
			Name:        "search",
			Flags:       getFlagSearch(),
			Subcommands: subCommandBookSearch(),
		},
	}
}

func subCommandBookSearch() []cli.Command {
	return []cli.Command{
		{
			Name:   "query",
			Action: actionQuery,
		},
		{
			Name:   "tag",
			Action: actionTag,
		},
	}
}
func actionBookNumber(c *cli.Context) {
	if c.NArg() == 0 {
		rand.Seed(time.Now().UnixNano())
		number := rand.Intn(1000)
		getBookNumber(fmt.Sprintf(bookApi, strconv.Itoa(int(number))))
	}

}
func actionBookDetail(c *cli.Context) {
	if c.NArg() == 1 {
		getBookDetail(c.Args().Get(0))
	}
}
func actionQuery(c *cli.Context) {
	if c.NArg() == 1 {
		//fmt.Println(c.String("count"))
		url := fmt.Sprintf(bookSearchByQuery, c.Args().Get(0), strconv.Itoa(c.Int("count")))
		getBookSearch(url)
	}
}

func actionTag(c *cli.Context) {
	if c.NArg() == 1 {
		url := fmt.Sprintf(bookSearchByTag, c.Args().Get(0), c.Int("count"))
		getBookSearch(url)
	}

}
func getBookNumber(url string) {
	var allBook []objects.BookInfo
	data := utils.Response(url, "books")
	for _, one := range data.Array() {
		var oneBook objects.BookInfo
		oneBook.Title = one.Get("title").String()
		oneBook.Subtitle = one.Get("subtitle").String()
		oneBook.Isbn10 = one.Get("isbn10").String()
		oneBook.Isbn13 = one.Get("isbn13").String()
		oneBook.URL = one.Get("url").String()
		oneBook.Price = one.Get("price").String()
		allBook = append(allBook, oneBook)
	}
	AllData := objects.BookAll{
		BookCollection: allBook,
	}
	AllData.Template()
}

func getBookDetail(isbnNumber string) {
	var OneBookDetail objects.BookDetail
	data := utils.Response(fmt.Sprintf(bookDetail, isbnNumber), "")
	OneBookDetail.Average = data.Get("rating").Get("average").String()
	OneBookDetail.Author = data.Get("author").String()
	OneBookDetail.PublicationDate = data.Get("pubdate").String()
	OneBookDetail.Image = data.Get("image").String()
	OneBookDetail.ID = data.Get("id").String()
	OneBookDetail.Publisher = data.Get("publisher").String()
	OneBookDetail.Price = data.Get("price").String()
	OneBookDetail.Summary = data.Get("summary").String()
	OneBookDetail.URL = strings.TrimSpace(data.Get("url").String())
	OneBookDetail.Title = data.Get("title").String()
	OneBookDetail.Catalog = data.Get("catalog").String()
	OneBookDetail.Translator = data.Get("translator").String()
	OneBookDetail.Template()
}

func getBookSearch(url string) {
	getBookNumber(url)
}
