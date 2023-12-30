package app

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/ayberktandogan/melody/config"
	"github.com/posener/complete"
	"github.com/willabides/kongplete"
)

const help = `♪ Melody is a CLI application for Spotify Music.`

func Main() {
	cli := cliBase{}
	userConfig, err := LoadUserConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
	}

	kongOptions := []kong.Option{
		kong.UsageOnError(),
		kong.Description(help),
		kong.BindTo(cli, (*cliInterface)(nil)),
		kong.Bind(userConfig),
		kong.AutoGroup(func(parent kong.Visitable, flag *kong.Flag) *kong.Group {
			node, ok := parent.(*kong.Command)
			if !ok {
				return nil
			}
			return &kong.Group{
				Key:   node.Name,
				Title: "Command flags:",
			}
		}),
		kong.Vars{
			"version": config.Config.Version,
		},
		kong.HelpOptions{
			Compact:             true,
			NoExpandSubcommands: true,
		},
	}
	parser, err := kong.New(&cli, kongOptions...)
	if err != nil {
		panic(fmt.Errorf("%w", err))
	}

	ctx, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	kongplete.Complete(parser,
		kongplete.WithPredictor("dir", complete.PredictDirs("*")),
		kongplete.WithPredictor("hclfile", complete.PredictFiles("*.hcl")),
		kongplete.WithPredictor("file", complete.PredictFiles("*")),
	)

	ctx.Run()
}
