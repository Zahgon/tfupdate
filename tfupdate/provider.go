package tfupdate

import (
	"context"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// ProviderUpdater is a updater implementation which updates the provider version constraint.
type ProviderUpdater struct {
	name    string
	version string
}

// NewProviderUpdater is a factory method which returns an ProviderUpdater instance.
func NewProviderUpdater(name string, version string) (Updater, error) {
	_ = "STUB: not implemented"
	return *new(Updater), nil
}

// Update updates the provider version constraint.
// Note that this method will rewrite the AST passed as an argument.
func (u *ProviderUpdater) Update(_ context.Context, mc *ModuleContext, filename string, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// skip a lock file.

func (u *ProviderUpdater) updateTerraformBlock(mc *ModuleContext, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// If the name contains /, assume that a namespace is intended and check the source.

// The hclwrite.Attribute doesn't have enough AST for object type to check.
// Get the attribute as a native hcl.Attribute as a compromise.

// There are some variations on the syntax of required_providers.
// So we check a type of the value and switch implementations.
// If the expression can be parsed as a static expression and it's type is a primitive,
// then it's a legacy string syntax.

// Otherwise, it's an object syntax.

func (u *ProviderUpdater) updateTerraformRequiredProvidersBlockAsObject(p *hclwrite.Block, name string, hclAttr *hcl.Attribute) error {
	_ = "STUB: not implemented"
	// terraform {
	//   required_providers {
	//     aws = {
	//       source  = "hashicorp/aws"
	//       version = "2.65.0"
	//
	//       configuration_aliases = [
	//         aws.primary,
	//         aws.secondary,
	//       ]
	//     }
	//   }
	// }
	return nil
}

// If the version key is missing, just ignore it.

// Updating the whole object loses original sort order and comments.
// At the time of writing, there is no way to update a value inside an
// object directly while preserving original tokens.
//
// Since we fully understand the valid syntax, we compromise and read the
// tokens in order, updating the bytes directly.
// It's apparently a fragile dirty hack, but I didn't come up with the better
// way to do this.

// find key of version
// Although not explicitly stated in the required_providers documentation,
// a TokenQuotedLit is also valid token. Strict speaking there are more
// variants because the left hand side of object key accepts an expression in
// HCL. For accurate implementation, it should be implemented using the
// original parser.
// nolint:staticcheck // QF1001: could apply De Morgan's law
// Ignore it in favor of subjective readability.

// find =

// find value of old version
// nolint:staticcheck // QF1001: could apply De Morgan's law
// Ignore it in favor of subjective readability.

// Since I've checked for the existence of the version key in advance,
// if we reach here, we found the token to be updated.
// So we now update bytes of the token in place.

// detectVersionInObject parses an object expression and detects a value for
// the "version" key.
// If the version key is missing, just returns an empty string without an error.
func detectVersionInObject(hclAttr *hcl.Attribute) (string, error) {
	_ = "STUB: not implemented"
	// The configuration_aliases syntax isn't directly related version updateing,
	// but it contains provider references and causes an parse error without an EvalContext.
	// So we treat the expression as a hcl.ExprMap to avoid fully decoding the object.
	return "", nil
}

func (u *ProviderUpdater) updateTerraformRequiredProvidersBlockAsString(p *hclwrite.Block) {
	_ = "STUB: not implemented"
	// terraform {
	//   required_providers {
	//     aws = "2.65.0"
	//   }
	// }
	return
}

func (u *ProviderUpdater) updateProviderBlock(f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// set a version to attribute value only if the key exists
