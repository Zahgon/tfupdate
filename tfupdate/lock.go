package tfupdate

import (
	"context"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/minamijoyo/tfupdate/lock"
	"github.com/minamijoyo/tfupdate/tfregistry"
)

// LockUpdater is a updater implementation which updates the dependency lock file.
type LockUpdater struct {
	platforms []string

	// index is a cached index for updating dependency lock files.
	index lock.Index

	// tfregistryConfig is a configuration for Terraform Registry API.
	tfregistryConfig tfregistry.Config
}

// NewLockUpdater is a factory method which returns an LockUpdater instance.
func NewLockUpdater(platforms []string, tfregistryConfig tfregistry.Config) (Updater, error) {
	_ = "STUB: not implemented"
	// Create a new index with the provided registry config
	return *new(Updater), nil
}

// Update updates the dependency lock file.
// Note that this method will rewrite the AST passed as an argument.
func (u *LockUpdater) Update(ctx context.Context, mc *ModuleContext, filename string, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip other than the lock file.

// updateLockfile updates the dependency lock file.
func (u *LockUpdater) updateLockfile(ctx context.Context, mc *ModuleContext, f *hclwrite.File) error {
	_ = "STUB: not implemented"
	return nil
}

// Unsupported formats, such as legacy abbreviated notation, will result
// in parse errors, but should be ignored without returning an error if
// possible.

// update the existing provider block

// create a new provider block

// updateProviderBlock updates the provider block in the dependency lock file.
func (u *LockUpdater) updateProviderBlock(ctx context.Context, pBlock *hclwrite.Block, p SelectedProvider) error {
	_ = "STUB: not implemented"
	return nil
}

// a version attribute found

// Avoid unnecessary recalculations if no version change

//Strictly speaking, constraints can contain multiple constraint expressions,
//including comparison operators, but in the tfupdate use case, we assume
//that the required_providers are pinned to a specific version to detect the
//required version without terraform init, so we can simply specify the
//constraints attribute as the same as the version. This may differ from what
//terraform generates, but we expect that it doesn't matter in practice.

// Calculate the hash value of the provider.
// Note that the provider will be downloaded if cache miss.

// fullyQualifiedProviderAddress converts the short form of the provider
// address into the fully qualified form.
// Example: hashicorp/null => registry.terraform.io/hashicorp/null
// If BaseURL is set (e.g., https://registry.opentofu.org/), it will use its hostname
// instead of the default one (e.g., hashicorp/null => registry.opentofu.org/hashicorp/null).
func (u *LockUpdater) fullyQualifiedProviderAddress(address string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Since .terraform.lock.hcl was introduced from v0.14, we assume that
// provider address is qualified with namespaces at least. We won't support
// implicit legacy things.

// If BaseURL is set, use its hostname

// Use the hostname from BaseURL with type casting to svchost.Hostname
