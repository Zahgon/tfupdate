package command

import (
	"github.com/minamijoyo/tfupdate/release"
	"github.com/mitchellh/cli"
	"github.com/spf13/afero"
)

// Meta are the meta-options that are available on all or most commands.
type Meta struct {
	// UI is a user interface representing input and output.
	UI cli.Ui

	// Fs is an afero filesystem.
	Fs afero.Fs
}

// newRelease is a factory method which returns an Release implementation.
func newRelease(sourceType string, source string) (release.Release, error) {
	_ = "STUB: not implemented"
	return *new(release.Release), nil
}
