package tfupdate

import (
	"context"
	"regexp"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

// moduleSourceRegexp is a regular expression for module source.
// This is not a complete module source definition, but it is sufficient to
// parse version. Note that a git reference can be branch name, so we need to
// check if it seems to be a version number.
// https://www.terraform.io/docs/modules/sources.html
var moduleSourceRegexp = regexp.MustCompile(`(.+)\?ref=v([0-9]+(\.[0-9]+)*(-.*)*)`)

// ModuleUpdater is a updater implementation which updates the module version constraint.
type ModuleUpdater struct {
	name      string
	nameRegex *regexp.Regexp
	version   string
}

// NewModuleUpdater is a factory method which returns an ModuleUpdater instance.
func NewModuleUpdater(name string, version string, nameRegex *regexp.Regexp) (Updater, error) {
	_ = "STUB: not implemented"
	return *new(Updater), nil
}

// Update updates the module version constraint.
// Note that this method will rewrite the AST passed as an argument.
func (u *ModuleUpdater) Update(_ context.Context, _ *ModuleContext, filename string, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// skip a lock file.

func (u *ModuleUpdater) match(name string) bool { _ = "STUB: not implemented"; return false }

func (u *ModuleUpdater) updateModuleBlock(f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// If this module is a target module

// The source attribute doesn't have a version number.
// Set a version to attribute value only if the version key exists.

// The source attribute has a version number.
// Update a version reference in the source value.

// parseModuleSource parses module source and returns module name and version.
func parseModuleSource(a *hclwrite.Attribute) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// no version number
