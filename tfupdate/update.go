package tfupdate

import (
	"context"
	"io"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// Updater is an interface which updates a version constraint in HCL.
type Updater interface {
	// Update updates a version constraint.
	// Note that this method will rewrite the AST passed as an argument.
	Update(ctx context.Context, mc *ModuleContext, filename string, f *hclwrite.File) error
}

// NewUpdater is a factory method which returns an Updater implementation.
func NewUpdater(o Option) (Updater, error) { _ = "STUB: not implemented"; return *new(Updater), nil }

// UpdateHCL reads HCL from io.Reader, updates version constraints
// and writes updated contents to io.Writer.
// If contents changed successfully, it returns true, or otherwise returns false.
// If an error occurs, Nothing is written to the output stream.
func UpdateHCL(ctx context.Context, mc *ModuleContext, r io.Reader, w io.Writer, filename string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// safeParseConfig parses config and recovers if panic occurs.
// The current hclwrite implementation is no perfect and will panic if
// unparseable input is given. We just treat it as a parse error so as not to
// surprise users of tfupdate.
func safeParseConfig(src []byte, filename string, start hcl.Pos) (f *hclwrite.File, e error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set a return value from panic recover
