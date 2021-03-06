package arbortools

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	arborbuild "github.com/arborlang/ArborGo/lib"
	"github.com/urfave/cli"
)

// Build is the entrypoint for building
type Build struct{}

// GetName returns the name of the command
func (b Build) GetName() string { return "build" }

// Category returns the catagory
func (b Build) Category() string { return "Build" }

// Action builds the project
func (b Build) Action(c *cli.Context) {
	if len(c.Args()) == 0 {
		fmt.Println("No file to build!")
		os.Exit(-1)
	}
	build := c.Args()[0]
	out := c.String("output")
	if out == "__FILE__" {
		fn := filepath.Base(build)
		out = strings.TrimSuffix(fn, path.Ext(fn))
	}
	fout, err := os.Create(out)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fin, err := os.Open(build)
	if err = arborbuild.BuildToWast(fin, fout); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// Flags returns the Flags for the command
func (b Build) Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  "wast",
			Usage: "Compile to WebAssembly text format",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "The file to write the compile file to",
			Value: "__FILE__",
		},
	}
}

// Help Describe the command
func (b Build) Help() map[string]string {
	mp := map[string]string{}
	mp["description"] = "Build an arbor project into an executable"
	mp["usage"] = "arbor build <options> [files]"
	return mp
}
