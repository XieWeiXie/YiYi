package objects

import (
	"fmt"
	"html/template"
	"os"
)

type ContentInfo struct {
	Author
	ID              string
	Title           string
	Auth            string
	AuthorIntroduce string
	MakeTime        string
	LastUpdateDate  string
	WebURL          string
	GuideWord       string
	Audio           string
	Anchor          string
	EditorEmail     string
}

type ContentPassage struct {
	Content string
}

type Author struct {
	UserID    string
	UserName  string
	Desc      string
	WbName    string
	Summary   string
	FansTotal string
}

func (c ContentInfo) Template() {
	t := template.New("New Template for ContentInfo")
	t, _ = t.Parse(`
Show ContentInfo from one App:
    ID: {{.ID}}
    Title: {{.Title}}
    Auth: {{.Auth}}
    AuthorIntroduce: {{.AuthorIntroduce}}
    MakeTime: {{.MakeTime}}
    LastUpdateDate: {{.LastUpdateDate}}
    WebURL: {{.WebURL}}
    GuideWord: {{.GuideWord}}
    Audio: {{.Audio}}
    Anchor: {{.Anchor}}
    EditorEmail: {{.EditorEmail}}
    AuthorInfo:
       UserID: {{.UserID}}
       UserName: {{.UserName}}
       Desc: {{.Desc}}
       WbName: {{.WbName}}
       Summary: {{.Summary}}
       FansTotal: {{.FansTotal}}

`)
	t.Execute(os.Stdout, c)
}

type ShareInfo struct {
	ShareList []map[string]ShareDetail
}

type ShareDetail struct {
	Title    string
	Desc     string
	Link     string
	ImageURL string
}

func (s ShareInfo) String() string {
	return fmt.Sprintf(`
Show ShareInfo from One App:
    wx:
        Title: %s,
        Desc: %s,
        Link: %s,
        ImgURL: %s,

    wx_timeline:
        Title: %s,
        Desc: %s,
        Link: %s,
        ImgURL: %s, 

    weibo:
        Title: %s,
        Desc: %s,
        Link: %s,
        ImgURL: %s,

    qq:
        Title: %s,
        Desc: %s,
        Link: %s,
        ImgURL: %s, 
`, s.ShareList[0]["wx"].Title, s.ShareList[0]["wx"].Desc, s.ShareList[0]["wx"].Link, s.ShareList[0]["wx"].ImageURL,
		s.ShareList[1]["wx_timeline"].Title, s.ShareList[1]["wx_timeline"].Desc, s.ShareList[1]["wx_timeline"].Link, s.ShareList[1]["wx_timeline"].ImageURL,
		s.ShareList[2]["weibo"].Title, s.ShareList[1]["weibo"].Desc, s.ShareList[1]["weibo"].Link, s.ShareList[1]["weibo"].ImageURL,
		s.ShareList[3]["qq"].Title, s.ShareList[3]["qq"].Desc, s.ShareList[3]["qq"].Link, s.ShareList[3]["qq"].ImageURL)
}
