package tfupdate

import (
	"context"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

// TerraformUpdater is a updater implementation which updates the terraform version constraint.
type TerraformUpdater struct {
	version string
}

// NewTerraformUpdater is a factory method which returns an TerraformUpdater instance.
func NewTerraformUpdater(version string) (Updater, error) {
	_ = "STUB: not implemented"
	return *new(Updater), nil
}

// Update updates the terraform version constraint.
// Note that this method will rewrite the AST passed as an argument.
func (u *TerraformUpdater) Update(_ context.Context, _ *ModuleContext, filename string, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// skip a lock file.

// set a version to attribute value only if the key exists
