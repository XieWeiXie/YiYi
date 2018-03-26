package objects

import (
	"html/template"
	"os"
)

type MovieAll struct {
	MovieCollection []MovieInfo
}

type MovieInfo struct {
	ID            string
	Title         string
	OriginalTitle string
	CollectCount  string
	Stars         string
	Average       string
	Casts         []Cast
}
type Cast struct {
	Alt  string
	Name string
	ID   string
}

func (m MovieAll) Template() {
	t := template.New("New template for MovieAll")
	t, _ = t.Parse(`
Show Movie Info from DouBan Api.
    MovieCollections In theaters:
    {{range .MovieCollection}}
        ID: {{.ID}}
        Title: {{.Title}}
        OriginalTitle: {{.OriginalTitle}}
        Average: {{.Average}}
        CollectionCount: {{.CollectCount}}
        Stars: {{.Stars}}
        Casts:
              {{range .Casts}}
              Alt: {{.Alt}}
              Name: {{.Name}}
              ID: {{.ID}}
              {{end}}
    {{end}}
`)
	t.Execute(os.Stdout, m)
}

type MovieDetail struct {
	Average       string
	ReviewsCount  string
	WishCount     string
	Year          string
	ID            string
	MobileURL     string
	Title         string
	ShareURL      string
	Countries     string
	Genres        string
	CollectCount  string
	Casts         []Cast
	OriginalTitle string
	Directors     string
	CommentsCount string
	RatingsCount  string
	Aka           string
	Summary       string
}

func (m MovieDetail) Template() {
	t := template.New("New Template for MovieDetail")
	t, _ = t.Parse(`
Show Get Movie In theaters from DouBan Api:
    ID: {{.ID}}
    Title: {{.Title}}
    OriginalTitle: {{.OriginalTitle}}
    Directors: {{.Directors}}
    Average: {{.Average}}
    ReviewsCount: {{.ReviewsCount}}
    WishCount: {{.WishCount}}
    CommentsCount: {{.CommentsCount}}
    RatingsCount: {{.RatingsCount}}
    Aka: {{.Aka}}
    Year: {{.Year}}
    MobileURL: {{.MobileURL}}
    ShareURL: {{.ShareURL}}
    Countries: {{.Countries}}
    Genres: {{.Genres}}
    CollectCount: {{.CollectCount}}
    Casts:
        {{range .Casts}}
        Alt: {{.Alt}}
        Name: {{.Name}}
        ID: {{.ID}}
        {{end}}
    Summary: {{.Summary}}
`)
	t.Execute(os.Stdout, m)
}
