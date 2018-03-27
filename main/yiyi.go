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
	Usage       = "Reader [global options] command [command options] [arguments...]"
	UsageText   = "Reader [global options] command [command options] [arguments...]"
	Version     = `
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

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Description = Description
	app.Usage = Usage
	app.UsageText = Usage
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.CommandNotFound = func(context *cli.Context, command string) {
		fmt.Printf("[[WARNING] Not Found Command: %s\n", command)
		fmt.Printf("[MESSAGE] Please Type: Reader --help")
	}
	app.Commands = reader.Commands()
	app.Run(os.Args)

}
