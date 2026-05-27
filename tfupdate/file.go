package tfupdate

import (
	"context"
)

// UpdateFile updates version constraints in a single file.
// We use an afero filesystem here for testing.
func UpdateFile(ctx context.Context, mc *ModuleContext, filename string) error {
	_ = "STUB: not implemented"
	return nil
}

// Write contents back to source file if changed.

// We should be able to choose whether to format output or not.
// However, the current implementation of (*hclwrite.Body).SetAttributeValue()
// does not seem to preserve an original SpaceBefore value of attribute.
// So, we need to format output here.

// UpdateDir updates version constraints for files in a given directory.
// If a recursive flag is true, it checks and updates recursively.
// skip hidden directories such as .terraform or .git.
// It also skips unsupported file type.
func UpdateDir(ctx context.Context, current *ModuleContext, dirname string) error {
	_ = "STUB: not implemented"
	return nil
}

// if a path of entry matches ignorePaths, skip it.

// if an entry is a directory

// skip directory if a recursive flag is false

// skip hidden directories such as .terraform or .git

// if an entry is a file
// nolint:staticcheck // QF1001: could apply De Morgan's law
// Ignore it in favor of subjective readability.

// skip unsupported file type

// UpdateFileOrDir updates version constraints in a given file or directory.
func UpdateFileOrDir(ctx context.Context, gc *GlobalContext, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// if an entry is a directory

// if an entry is a file
// Note that even if only the filename is specified, the directory containing
// it is read for module context analysis.

// When the filename is intentionally specified,
// we should not ignore it by its extension as much as possible.
