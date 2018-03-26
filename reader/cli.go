package reader

import (
	"YiYi/commands/douban"
	"YiYi/commands/one"

	"github.com/urfave/cli"
)

func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:        "book",
			Usage:       douban.BookUsage,
			Subcommands: douban.SubBookCommand(),
		},
		{
			Name:        "story",
			Usage:       one.StoryUsage,
			Subcommands: one.SubStoryCommand(),
		},
		{
			Name:        "movie",
			Usage:       douban.MovieUsage,
			Subcommands: douban.SubMovieCommand(),
		},
	}
}
