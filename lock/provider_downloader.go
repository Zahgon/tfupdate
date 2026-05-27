package lock

import (
	"context"
	"net/http"

	"github.com/minamijoyo/tfupdate/tfregistry"
)

// PackageDownloaderAPI is an interface for downloading provider package.
// Provider packages are downloaded from the HashiCorp release server,
// GitHub release page or somewhere else.
// Therefore we distinct this API from the Terraform Registry API.
// The API specification is not documented.
type ProviderDownloaderAPI interface {
	// ProviderDownload downloads a provider package.
	ProviderDownload(ctx context.Context, req *ProviderDownloadRequest) (*ProviderDownloadResponse, error)
}

// ProviderDownloaderClient implements the ProviderDownloaderAPI interface
type ProviderDownloaderClient struct {
	// api is an instance of tfregistry.API interface.
	// It can be replaced for testing.
	api tfregistry.API

	// httpClient is a http client which communicates with the ProviderDownloaderAPI.
	httpClient *http.Client
}

// NewProviderDownloaderClient is a factory method which returns a ProviderDownloaderClient instance.
func NewProviderDownloaderClient(config tfregistry.Config) (*ProviderDownloaderClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProviderDownloadRequest is a request type for ProviderDownload.
type ProviderDownloadRequest struct {
	// (required): the namespace portion of the address of the requested provider.
	Namespace string `json:"namespace"`
	// (required): the type portion of the address of the requested provider.
	Type string `json:"type"`
	// (required): the version selected to download.
	Version string `json:"version"`
	// (required): a keyword identifying the operating system that the returned package should be compatible with, like "linux" or "darwin".
	OS string `json:"os"`
	// (required): a keyword identifying the CPU architecture that the returned package should be compatible with, like "amd64" or "arm".
	Arch string `json:"arch"`
}

// ProviderDownloadResponse is a response type for ProviderDownload.
type ProviderDownloadResponse struct {
	// filename is the filename for zipData.
	filename string

	// zipData is the raw byte sequence of the provider package.
	zipData []byte

	// shaSumsData is the raw byte sequence of the provider shasum file.
	shaSumsData []byte
}

// ProviderDownload downloads a provider package.
func (c *ProviderDownloaderClient) ProviderDownload(ctx context.Context, req *ProviderDownloadRequest) (*ProviderDownloadResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// download is a helper function that downloads contents from a given url.
func (c *ProviderDownloaderClient) download(ctx context.Context, url string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateSHA256Sum calculates the sha256 sum of the given byte sequence and
// checks whether it matches the expected hash value.
// The hash value is specified as a hexadecimal string.
func validateSHA256Sum(b []byte, sha256sum string) error { _ = "STUB: not implemented"; return nil }

// sha256sumAsHexString calculates the sha256 sum of the given byte sequence and
// returns it as a hexadecimal string.
func sha256sumAsHexString(b []byte) string { _ = "STUB: not implemented"; return "" }

// validateSHASumsData checks whether the SHA256Sum document contains a matching hash value for a given filename.
func validateSHASumsData(b []byte, filename string, sha256sum string) error {
	_ = "STUB: not implemented"
	return nil
}

// We expect that blank lines are not normally included, but to make the
// test data easier to read, ignore blank lines.

// Split rows into columns with spaces, but note that there are two spaces between the columns.
// e4453fbebf90c53ca3323a92e7ca0f9961427d2f0ce0d2b65523cc04d5d999c2  terraform-provider-null_3.2.1_darwin_arm64.zip

// ok

// not found
