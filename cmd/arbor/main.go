package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"plugin"

	arbor "github.com/arborlang/arbor-dev"
	"github.com/arborlang/arbor/internal"
	"github.com/urfave/cli"
)

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// LoadPlugins load the plugins for the tool chain
func LoadPlugins() []arbor.Command {
	plugs := []arbor.Command{
		arbortools.Build{},
		arbortools.Run{},
	}
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir = path.Join(dir, "plugins")
	doesExsit, err := exists(dir)
	if err != nil {
		panic(err)
	}
	if !doesExsit {
		return plugs
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			plug, err := plugin.Open(path)
			if err != nil {
				return err
			}
			cmdPlug, err := plug.Lookup("Command")
			if err != nil {
				return err
			}
			cmd, ok := cmdPlug.(arbor.Command)
			if !ok {
				return fmt.Errorf("plugin %s doesn't implement the plugins.Command interface", path)
			}
			plugs = append(plugs, cmd)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return plugs
}

// Version is the the toolchain version
var Version = "0.0.0-rc0"

func main() {
	plugs := LoadPlugins()
	app := cli.NewApp()
	app.Flags = []cli.Flag{}
	app.Author = "Yoseph Radding"
	app.Description = "The Arbor Language Tool Chain"
	app.Name = "The Tool chain to manage Arbor code"
	app.Version = Version
	cmds := []cli.Command{}
	for _, plug := range plugs {
		cmds = append(cmds, cli.Command{
			Name:        plug.GetName(),
			Action:      plug.Action,
			Description: plug.Help()["description"],
			UsageText:   plug.Help()["usage"],
			Flags:       plug.Flags(),
		})
	}
	app.Commands = cmds
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
