package tfupdate

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// allMatchingBlocks returns all matching blocks from the body that have the
// given name and labels or returns an empty list if there is currently no
// matching block.
func allMatchingBlocks(b *hclwrite.Body, typeName string, labels []string) []*hclwrite.Block {
	_ = "STUB: not implemented"
	return nil
}

// allMatchingBlocksByType returns all matching blocks from the body that have the
// given name or returns an empty list if there is currently no matching block.
// This method is useful when you want to ignore label differences.
func allMatchingBlocksByType(b *hclwrite.Body, typeName string) []*hclwrite.Block {
	_ = "STUB: not implemented"
	return nil
}

// getHCLNativeAttribute gets hclwrite.Attribute as a native hcl.Attribute.
// At the time of writing, there is no way to do with the hclwrite AST,
// so we build low-level byte sequences and parse an attribute as a
// hcl.Attribute on memory.
// If not found, returns nil without an error.
func getHCLNativeAttribute(body *hclwrite.Body, name string) (*hcl.Attribute, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build low-level byte sequences

// parse an expression as a hcl.File.
// Note that an attribute may contains references, which are defined outside the file.
// So we cannot simply use hclsyntax.ParseExpression or hclsyntax.ParseConfig here.
// We need to use a loe-level parser not to resolve all references.

// getAttributeValueAsString returns a value of Attribute as string.
// There is no way to get value as string directly,
// so we parses tokens of Attribute and build string representation.
// The returned value is unquoted.
func getAttributeValueAsUnquotedString(attr *hclwrite.Attribute) string {
	_ = "STUB: not implemented"
	// find TokenEqual
	return ""
}

// TokenIdent records SpaceBefore, but we should ignore it here.

// unquote

// tokensForListPerLine builds a hclwrite.Tokens for a given list, but breaks the line for each element.
func tokensForListPerLine(list []string) hclwrite.Tokens {
	_ = "STUB: not implemented"
	// The original TokensForValue implementation does not break line by line for list,
	// so we build a token sequence by ourselves.
	return *new(hclwrite.Tokens)
}
