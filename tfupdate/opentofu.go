package tfupdate

import (
	"context"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

// OpenTofuUpdater is a updater implementation which updates the OpenTofu version constraint.
type OpenTofuUpdater struct {
	version string
}

// NewOpenTofuUpdater is a factory method which returns an OpenTofuUpdater instance.
func NewOpenTofuUpdater(version string) (Updater, error) {
	_ = "STUB: not implemented"
	return *new(Updater), nil
}

// Update updates the OpenTofu version constraint.
// Note that this method will rewrite the AST passed as an argument.
func (u *OpenTofuUpdater) Update(_ context.Context, _ *ModuleContext, filename string, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// skip a lock file.

// set a version to attribute value only if the key exists
