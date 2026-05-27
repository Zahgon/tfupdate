package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

// Version is a version number.
var version = "0.9.4"

// UI is a user interface which is a global variable for mocking.
var UI cli.Ui

func init() {
	UI = &cli.BasicUi{
		Writer: os.Stdout,
	}
}

func main() {
	log.SetOutput(logOutput())

	// #nosec G706: Log injection via taint analysis
	log.Printf("[INFO] CLI args: %#v", os.Args)

	commands := initCommands()

	args := os.Args[1:]

	c := &cli.CLI{
		Name:                  "tfupdate",
		Version:               version,
		Args:                  args,
		Commands:              commands,
		HelpWriter:            os.Stdout,
		Autocomplete:          true,
		AutocompleteInstall:   "install-autocomplete",
		AutocompleteUninstall: "uninstall-autocomplete",
	}

	exitStatus, err := c.Run()
	if err != nil {
		UI.Error(fmt.Sprintf("Failed to execute CLI: %s", err))
	}

	os.Exit(exitStatus)
}

func logOutput() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

// default log writer is null device.

func initCommands() map[string]cli.CommandFactory { _ = "STUB: not implemented"; return nil }
