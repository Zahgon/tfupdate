package lock

import (
	"context"

	"github.com/minamijoyo/tfupdate/tfregistry"
)

// Index is an in-memory data store for caching provider hash values.
type Index interface {
	// GetOrCreateProviderVersion returns a cached provider version if available,
	// otherwise creates it.
	// address is a provider address such as hashicorp/null.
	// version is a version number such as 3.2.1.
	// platforms is a list of target platforms to generate hash values.
	// Target platform names consist of an operating system and a CPU architecture such as darwin_arm64.
	GetOrCreateProviderVersion(ctx context.Context, address string, version string, platforms []string) (*ProviderVersion, error)
}

// index is an implementation for Index interface.
type index struct {
	// providers is a dictionary of providerIndex.
	// The key is a provider address such as hashicorp/null.
	providers map[string]*providerIndex

	// papi is a ProviderDownloaderAPI interface implementation used for downloading provider.
	papi ProviderDownloaderAPI
}

// NewIndexFromConfig returns a new instance of Index with the given registry config.
func NewIndexFromConfig(config tfregistry.Config) (Index, error) {
	_ = "STUB: not implemented"
	return *new(Index), nil
}

// NewIndex returns a new instance of Index with the given provider downloader API.
func NewIndex(papi ProviderDownloaderAPI) Index { _ = "STUB: not implemented"; return *new(Index) }

// GetOrCreateProviderVersion returns a cached provider version if available,
// otherwise creates it.
func (i *index) GetOrCreateProviderVersion(ctx context.Context, address string, version string, platforms []string) (*ProviderVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cache miss

// Delegate to ProviderIndex.

// The providerIndex holds multiple version data for a specific provider.
type providerIndex struct {
	// address is a provider address such as hashicorp/null.
	address string

	// versions is a dictionary of ProviderVersion.
	// The key is a version number such as 3.2.1.
	versions map[string]*ProviderVersion

	// papi is a ProviderDownloaderAPI interface implementation used for downloading provider.
	papi ProviderDownloaderAPI
}

// newProviderIndex returns a new instance of providerIndex.
func newProviderIndex(address string, papi ProviderDownloaderAPI) *providerIndex {
	_ = "STUB: not implemented"
	return nil
}

// getOrCreateProviderVersion returns a cached provider version if available,
// otherwise creates it.
func (pi *providerIndex) getOrCreateProviderVersion(ctx context.Context, version string, platforms []string) (*ProviderVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cache miss

// createProviderVersion downloads the specified provider, calculates the hash
// value and returns an instance of the ProviderVersion.
func (pi *providerIndex) createProviderVersion(ctx context.Context, version string, platforms []string) (*ProviderVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Download a given provider from registry.

// Currently the Terraform Registry returns the zh hash for all platforms,
// but not the h1 hash, so the h1 hash has to be calculated separately.
// We need to calculate the values for each platform and merge the results.

// newProviderDownloadRequest is a helper function for building the parameters for downloading provider.
// address is a provider address such as hashicorp/null.
// version is a version number such as 3.2.1.
// platform is a target platform name such as darwin_arm64.
func newProviderDownloadRequest(address string, version string, platform string) (*ProviderDownloadRequest, error) {
	_ = "STUB: not implemented"
	// We parse an provider address by using the terraform-registry-address
	// library to support fully qualified addresses such as
	// registry.terraform.io/hashicorp/null in the future, but note that the
	// current ProviderDownloaderClient implementation only supports the public
	// standard registry (registry.terraform.io).
	return nil, nil
}

// Since .terraform.lock.hcl was introduced from v0.14, we assume that
// provider address is qualified with namespaces at least. We won't support
// implicit legacy things.

// buildProviderVersion calculates hash values from the ProviderDownloadResponse
// and returns an instance of the ProviderVersion.
func buildProviderVersion(address string, version string, platform string, res *ProviderDownloadResponse) (*ProviderVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
