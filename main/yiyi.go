package main

import (
	"YiYi/reader"
	"fmt"

	"os"

	"github.com/urfave/cli"
)

const (
	Name        = "YiYi"
	Description = "An application for book, movie, and story from DouBan and One App."
	//Usage       = "Reader [global options] command [command options] [arguments...]"
	UsageText = "YiYi [global options] command [command options] [arguments...]"
	Version   = `
    ___       ___       ___       ___   
   /\__\     /\  \     /\__\     /\  \  
  |::L__L   _\:\  \   |::L__L   _\:\  \ 
  |:::\__\ /\/::\__\  |:::\__\ /\/::\__\
  /:;;/__/ \::/\/__/  /:;;/__/ \::/\/__/
  \/__/     \:\__\    \/__/     \:\__\  
             \/__/               \/__/   v0.0.1
`
	Author = "xieWei"
	Email  = "wuxiaoshen@shu.edu.cn"
)

var AppHelpTemplate = `YiYi is a tool for reading with DouBan and One APP api.


NAME:
   {{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

VERSION:
   {{.Version}}{{end}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if len .Authors}}

AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

COMMANDS:{{range .VisibleCategories}}{{if .Name}}

   {{.Name}}:{{end}}{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

GLOBAL OPTIONS:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

COPYRIGHT:
   {{.Copyright}}{{end}}
`

func main() {
	app := cli.NewApp()
	app.CustomAppHelpTemplate = AppHelpTemplate
	app.Name = Name
	app.Description = Description
	app.Usage = Description
	app.UsageText = UsageText
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.CommandNotFound = func(context *cli.Context, command string) {
		fmt.Printf("[WARNING] Not Found Command: %s\n", command)
		fmt.Printf("[MESSAGE] Please Type: YiYi --help")
	}
	app.Commands = reader.Commands()
	app.Run(os.Args)

}
