package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/rwxrob/cmdtab"
)

var (
	err error
)

func init() {
	x := cmdtab.New("open")
	x.Version = `v1.0.0`
	x.Summary = `open a resource`
	x.Usage = `[<resource>]`
	x.Description = `The *open* subcommand aims to solve the issue of
	having multiple ways to open a resource (file, URL, ...). It detects
	the current platform and runs the default program for <resource>.`
	x.License = `MPL-2.0 License`
	x.Method = func(args []string) error {
		if len(args) == 0 {
			return x.UsageError()
		}

		switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", args[0]).Run()
		case "windows":
			err = exec.Command("start", args[0]).Run()
		case "darwin":
			err = exec.Command("open", args[0).Run()
		default:
			err = fmt.Errorf("platform not supported ...\n")
		}
		return err
	}
}
