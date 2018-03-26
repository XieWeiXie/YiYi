package objects

import (
	"fmt"
	"html/template"
	"os"
)

type Weather struct {
	CityName      string
	Date          string
	Temperature   string
	Humidity      string
	Climate       string
	WindDirection string
	Hurricane     string
}

func (W Weather) String() string {
	return fmt.Sprintf(`Weather Info for One App:
Informatations:
    CityName:      %s,
    Date:          %s,
    Temperature:   %s,
    Humidity:      %s,
    Climate:       %s,
    WindDirection: %s,
    Hurricane:     %s
`, W.CityName, W.Date, W.Temperature, W.Humidity, W.Climate, W.WindDirection, W.Hurricane)
}

type Menu struct {
	Vol  string
	List []ListInfo
}

type ListInfo struct {
	ContentType string
	ContentID   string
	Title       string
	Tag         map[string]string
}

func MapDealWith(args map[string]string) string {
	var returnValue string
	for _, value := range args {
		returnValue += value + " "
	}
	return returnValue

}
func (M Menu) Template() {
	t := template.New("New Template for Menu")
	t = t.Funcs(template.FuncMap{"mapDealWith": MapDealWith})
	t, _ = t.Parse(`
Show One List Menu form One App:
    Vol: {{.Vol}}
    {{if .List}}{{range .List}}
    ContentType: {{.ContentType}}
    ContentID:   {{.ContentID}}
    Title:       {{.Title}}
    Tag:         {{.Tag|mapDealWith}}
    {{end}}
    {{end}}
`)
	t.Execute(os.Stdout, M)
}
