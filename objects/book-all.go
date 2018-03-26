package objects

import (
	"html/template"
	"os"
)

type BookAll struct {
	BookCollection []BookInfo
}
type BookInfo struct {
	Title    string
	Subtitle string
	URL      string
	Isbn10   string
	Isbn13   string
	Price    string
}

func (b BookAll) Template() {
	t := template.New("New Template for book")
	t, _ = t.Parse(`
Show Book from DouBan Api:
    AllCollections:
	    {{range .BookCollection}}
	    Title: {{.Title}}
	    Subtitle: {{.Subtitle}}
	    URL: {{.URL}}
	    Isbn10: {{.Isbn10}}
	    Isbn13: {{.Isbn13}}
	    Price: {{.Price}}
	    {{end}}
`)
	t.Execute(os.Stdout, b)
}

type BookDetail struct {
	ID              string
	Title           string
	Average         string
	Author          string
	Publisher       string
	PublicationDate string
	Image           string
	Translator      string
	URL             string
	Summary         string
	Catalog         string
	Price           string
}

func (b BookDetail) Template() {
	t := template.New("New Template for BookDetail")
	t, _ = t.Parse(`
Show BookDetail for DonBan Api.
    ID: {{.ID}}
    Title: {{.Title}}
    Average: {{.Average}}
    Author: {{.Author}}
    Publisher: {{.Publisher}}
    PublicationDate: {{.PublicationDate}}
    ImageURL: {{.Image}}
    Translator: {{.Translator}}
    URL: {{.URL }}
      Summary: {{.Summary}}
    Catalog: {{.Catalog}}
    Price: {{.Price}}
`)
	t.Execute(os.Stdout, b)
}
