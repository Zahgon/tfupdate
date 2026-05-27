package lock

import (
	"context"
	"net/http"
	"net/url"

	"github.com/minamijoyo/tfupdate/tfregistry"
)

// mockTFRegistryClient is a mock implementation of tfregistry.API
type mockTFRegistryClient struct {
	metadataRes *tfregistry.ProviderPackageMetadataResponse
	err         error
}

var _ tfregistry.API = (*mockTFRegistryClient)(nil)

func (c *mockTFRegistryClient) ProviderPackageMetadata(_ context.Context, _ *tfregistry.ProviderPackageMetadataRequest) (*tfregistry.ProviderPackageMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *mockTFRegistryClient) ListModuleVersions(_ context.Context, _ *tfregistry.ListModuleVersionsRequest) (*tfregistry.ListModuleVersionsResponse, error) {
	_ = "STUB: not implemented"
	// dummy implementation as it's not used in tests
	return nil, nil
}

func (c *mockTFRegistryClient) ListProviderVersions(_ context.Context, _ *tfregistry.ListProviderVersionsRequest) (*tfregistry.ListProviderVersionsResponse, error) {
	_ = "STUB: not implemented"
	// dummy implementation as it's not used in tests
	return nil, nil
}

// newMockServer returns a new mock server for testing.
func newMockServer() (*http.ServeMux, *url.URL) { _ = "STUB: not implemented"; return nil, nil }

// newTestClient returns a new client for testing.
func newTestClient(mockServerURL *url.URL, config tfregistry.Config) *ProviderDownloaderClient {
	_ = "STUB: not implemented"
	return nil
}

// newMockZipData returns a new zip format data for testing.
func newMockZipData(filename string, contents string) ([]byte, error) {
	_ = "STUB: not implemented"
	// create a zip file in memory
	return nil, nil
}

// create a file in the zip file

// zip

// newMockShaSumsData returns a new shaSumsData for testing.
// To ensure that the dummy data can be re-used in other test cases, the
// function really creates a zip file in memory and calculates its sha256sum.
func newMockShaSumsData(name string, version string, platforms []string) ([]byte, error) {
	_ = "STUB: not implemented"
	// terraform-provider-dummy_v3.2.1_x5
	return nil, nil
}

// dummy_3.2.1_darwin_arm64

// create a zip file in memory.

// newMockProviderDownloadResponse returns a new ProviderDownloadResponse for testing.
func newMockProviderDownloadResponse(address string, version string, targetPlatform string, allPlatforms []string) (*ProviderDownloadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create a zip file in memory.

// create a valid dummy shaSumsData.

// newMockProviderDownloadResponses returns a new list of ProviderDownloadResponse for testing.
func newMockProviderDownloadResponses(address string, version string, targetPlatforms []string, allPlatforms []string) ([]*ProviderDownloadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMockIndex does not call the real API but returns preset mock provider version metadata.
func NewMockIndex(pvs []*ProviderVersion) Index { _ = "STUB: not implemented"; return *new(Index) }

// NewMockProviderVersion returns a mocked ProviderVersion for testing.
// This is actually a setter to all private fields, but should not be used
// except for generating test data from outside the package.
func NewMockProviderVersion(address string, version string, platforms []string, h1Hashes map[string]string, zhHashes map[string]string) *ProviderVersion {
	_ = "STUB: not implemented"
	return nil
}
