package douban

import (
	"YiYi/objects"
	"YiYi/utils"

	"fmt"

	"strconv"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
)

const (
	NowInTheaters   = "https://api.douban.com/v2/movie/in_theaters"
	MovieDetail     = "http://api.douban.com/v2/movie/subject/%s"
	ComingSoonMovie = "https://api.douban.com/v2/movie/coming_soon"
	MovieTop250     = "http://api.douban.com/v2/movie/top250"
)

var MovieUsage = "get movie info from DouBan API"

func SubMovieCommand() []cli.Command {
	return []cli.Command{
		{
			Name:   "now",
			Action: actionNowInTheaters,
		},
		{
			Name:   "detail",
			Action: actionDetail,
		},
		{
			Name:   "future",
			Action: actionFuture,
		},
		{
			Name:   "top",
			Action: actionTop,
		},
	}
}

func actionNowInTheaters(c *cli.Context) {
	if c.NArg() == 0 {
		getNowInTheaters(NowInTheaters)
	}
}

func actionDetail(c *cli.Context) {
	if c.NArg() == 1 {
		getOneMovieDetail(fmt.Sprintf(MovieDetail, c.Args().Get(0)))
	}
}

func actionFuture(c *cli.Context) {
	if c.NArg() == 1 {
		number, _ := strconv.Atoi(c.Args().Get(0))
		getMovieComingSoon(ComingSoonMovie, number)
	}
	if c.NArg() == 0 {
		getMovieComingSoon(ComingSoonMovie, 0)
	}
}

func actionTop(c *cli.Context) {
	if c.NArg() == 1 {
		number, _ := strconv.Atoi(c.Args().Get(0))
		getTopMovie(MovieTop250, number)
	}
	if c.NArg() == 0 {
		getTopMovie(MovieTop250, 0)
	}
}

func getNowInTheaters(url string) {
	var collectionMovie []objects.MovieInfo
	data := utils.Response(url, "subjects")
	for _, value := range data.Array() {
		var oneMovieInfo objects.MovieInfo
		var casts []objects.Cast
		for _, castValue := range value.Get("casts").Array() {
			var oneCast objects.Cast
			oneCast.ID = castValue.Get("id").String()
			oneCast.Name = castValue.Get("name").String()
			oneCast.Alt = castValue.Get("alt").String()
			casts = append(casts, oneCast)
		}
		oneMovieInfo.Title = value.Get("title").String()
		oneMovieInfo.Average = value.Get("rating.average").String()
		oneMovieInfo.Casts = casts
		oneMovieInfo.CollectCount = value.Get("collect_count").String()
		oneMovieInfo.OriginalTitle = value.Get("original_title").String()
		oneMovieInfo.Stars = value.Get("rating.stars").String()
		oneMovieInfo.ID = value.Get("id").String()
		collectionMovie = append(collectionMovie, oneMovieInfo)
	}
	all := objects.MovieAll{
		MovieCollection: collectionMovie,
	}
	all.Template()
}

func getOneMovieDetail(url string) {
	var movieDetail objects.MovieDetail
	data := utils.Response(url, "")
	movieDetail = oneMovie(data)
	movieDetail.Template()
}

func getMovieComingSoon(url string, number int) {
	data := utils.Response(url, "subjects")
	printOneMovie(data, number)

}

func getTopMovie(url string, number int) {
	data := utils.Response(url, "subjects")
	printOneMovie(data, number)
}

func oneMovie(args gjson.Result) objects.MovieDetail {
	var movieDetail objects.MovieDetail
	movieDetail.ID = args.Get("id").String()
	movieDetail.OriginalTitle = args.Get("original_title").String()
	movieDetail.CollectCount = args.Get("collect_count").String()
	movieDetail.Average = args.Get("rating.average").String()
	movieDetail.ReviewsCount = args.Get("reviews_count").String()
	movieDetail.WishCount = args.Get("wish_count").String()
	movieDetail.Year = args.Get("year").String()
	movieDetail.MobileURL = args.Get("mobile_url").String()
	movieDetail.Title = args.Get("title").String()
	movieDetail.ShareURL = args.Get("share_url").String()
	movieDetail.RatingsCount = args.Get("ratings_count").String()
	movieDetail.Summary = args.Get("summary").String()
	var allCasts []objects.Cast
	for _, castValue := range args.Get("casts").Array() {
		var oneValue objects.Cast
		oneValue.ID = castValue.Get("id").String()
		oneValue.Alt = castValue.Get("alt").String()
		oneValue.Name = castValue.Get("name").String()
		allCasts = append(allCasts, oneValue)
	}
	movieDetail.Casts = allCasts
	var allDirectors []string
	for _, directorsValue := range args.Get("directors").Array() {
		allDirectors = append(allDirectors, directorsValue.Get("name").String())
	}
	movieDetail.Directors = utils.ListConvertString(allDirectors)
	var allCountries []string
	for _, countryValue := range args.Get("countries").Array() {
		allCountries = append(allCountries, countryValue.String())
	}
	var allGenres []string
	for _, genresValue := range args.Get("genres").Array() {
		allGenres = append(allGenres, genresValue.String())
	}
	movieDetail.Countries = utils.ListConvertString(allCountries)
	movieDetail.Genres = utils.ListConvertString(allGenres)
	var allAka []string
	for _, akaValue := range args.Get("aka").Array() {
		allAka = append(allAka, akaValue.String())
	}
	movieDetail.Aka = utils.ListConvertString(allAka)
	return movieDetail

}

func printOneMovie(args gjson.Result, number int) {
	if number == 0 {
		for _, one := range args.Array() {
			movieDetail := oneMovie(one)
			movieDetail.Template()
		}
	} else {
		movieDetail := oneMovie(args.Array()[number])
		movieDetail.Template()
	}
}
