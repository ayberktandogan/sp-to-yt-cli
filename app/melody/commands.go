package app

import (
	"github.com/alecthomas/kong"
)

type cliInterface interface {
	getVersion() kong.VersionFlag

	getLogin() loginCmd
	getUser() userCmd
}

type cliBase struct {
	User    userCmd          `cmd:"" group:"request types" name:"user" short:"u" help:"User related commands"`
	Login   loginCmd         `cmd:"" group:"auth" name:"login" short:"l" help:"Login to Spotify"`
	Version kong.VersionFlag `name:"version" group:"global" short:"v" help:"Show version"`
}

var _ cliInterface = &cliBase{}

func (c *cliBase) getLogin() loginCmd           { return c.Login }
func (c *cliBase) getVersion() kong.VersionFlag { return c.Version }
func (c *cliBase) getUser() userCmd             { return c.User }
