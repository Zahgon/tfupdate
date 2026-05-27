package tfupdate

import (
	"regexp"

	"github.com/minamijoyo/tfupdate/tfregistry"
)

// Option is a set of parameters to update.
type Option struct {
	// A type of updater. Valid values are as follows:
	// - terraform
	// - provider
	// - module
	// - lock
	updateType string

	// If an updateType is terraform, there is no meaning.
	// If an updateType is provider or module, Set a name of provider or module.
	name string

	// a new version constraint
	version string

	// platforms is a list of target platforms to generate hash values.
	// Target platform names consist of an operating system and a CPU
	// architecture such as darwin_arm64.
	platforms []string

	// If a recursive flag is true, it checks and updates directories recursively.
	recursive bool

	// An array of regular expression for paths to ignore.
	ignorePaths []*regexp.Regexp

	// This field stores the compiled RE2 regex from the provide name parameter.
	// In case the sourceMatchType is set to regex this field is used to match the name.
	// In case the provided sourceMatchType is full this field is nil.
	nameRegex *regexp.Regexp

	// tfregistryConfig is a configuration for Terraform Registry API.
	tfregistryConfig tfregistry.Config
}

// NewOption returns an option.
func NewOption(updateType string, name string, version string, platforms []string, recursive bool, ignorePaths []string, sourceMatchType string, tfregistryConfig tfregistry.Config) (Option, error) {
	_ = "STUB: not implemented"
	return *new(Option), nil
}

func nameRegex(updateType string, name string, sourceMatchType string) (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MatchIgnorePaths returns whether any of the ignore conditions are met.
func (o *Option) MatchIgnorePaths(path string) bool { _ = "STUB: not implemented"; return false }
