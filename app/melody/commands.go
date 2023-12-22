package app

import "github.com/alecthomas/kong"

type cliInterface interface {
	getVersion() kong.VersionFlag

	getLogin() loginCmd
}

type cliBase struct {
	Login   loginCmd         `cmd:"" name:"login" short:"l" help:"Login to Spotify"`
	Version kong.VersionFlag `name:"version" short:"v" help:"Show version"`
}

var _ cliInterface = &cliBase{}

func (c *cliBase) getLogin() loginCmd           { return c.Login }
func (c *cliBase) getVersion() kong.VersionFlag { return c.Version }
