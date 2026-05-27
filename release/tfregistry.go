package release

import (
	"context"

	"github.com/minamijoyo/tfupdate/tfregistry"
)

// TFRegistryModuleRelease is a release implementation which provides version information with TFRegistryModule Release.
type TFRegistryModuleRelease struct {
	// api is an instance of tfregistry.API interface.
	// It can be replaced for testing.
	api tfregistry.API

	// namespace is a user name which owns the module.
	namespace string

	// name is a name of the module.
	name string

	// provider is a name of the provider.
	provider string
}

var _ Release = (*TFRegistryModuleRelease)(nil)

// NewTFRegistryModuleRelease is a factory method which returns an TFRegistryModuleRelease instance.
func NewTFRegistryModuleRelease(source string, config tfregistry.Config) (Release, error) {
	_ = "STUB: not implemented"
	return *new(Release), nil
}

// ListReleases returns a list of unsorted all releases including pre-release.
func (r *TFRegistryModuleRelease) ListReleases(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract versions from the response

// TFRegistryProviderRelease is a release implementation which provides version information with TFRegistryProvider Release.
type TFRegistryProviderRelease struct {
	// api is an instance of tfregistry.API interface.
	// It can be replaced for testing.
	api tfregistry.API

	// The user or organization the provider is owned by.
	namespace string

	// The type name of the provider.
	providerType string
}

var _ Release = (*TFRegistryProviderRelease)(nil)

// NewTFRegistryProviderRelease is a factory method which returns an TFRegistryProviderRelease instance.
func NewTFRegistryProviderRelease(source string, config tfregistry.Config) (Release, error) {
	_ = "STUB: not implemented"
	return *new(Release), nil
}

// ListReleases returns a list of unsorted all releases including pre-release.
func (r *TFRegistryProviderRelease) ListReleases(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract versions from the response
