package tfregistry

import (
	"context"
)

// ProviderPackageMetadataRequest is a request parameter for ProviderPackageMetadata().
// https://developer.hashicorp.com/terraform/internals/provider-registry-protocol#find-a-provider-package
type ProviderPackageMetadataRequest struct {
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

// ProviderPackageMetadataResponse is a response data for ProviderPackageMetadata().
// There are other response fields, but we define only those we need here.
type ProviderPackageMetadataResponse struct {
	// (required): the filename for this provider's zip archive as recorded in the "shasums" document, so that Terraform CLI can determine which of the given checksums should be used for this specific package.
	Filename string `json:"filename"`
	// (required): a URL from which Terraform can retrieve the provider's zip archive. If this is a relative URL then it will be resolved relative to the URL that returned the containing JSON object.
	DownloadURL string `json:"download_url"`
	// (required): the SHA256 checksum for this provider's zip archive as recorded in the shasums document.
	SHASum string `json:"shasum"`
	// (required): a URL from which Terraform can retrieve a text document recording expected SHA256 checksums for this package and possibly other packages for the same provider version on other platforms.
	SHASumsURL string `json:"shasums_url"`
}

// ProviderPackageMetadata returns a package metadata of a provider.
// https://developer.hashicorp.com/terraform/internals/provider-registry-protocol#find-a-provider-package
func (c *Client) ProviderPackageMetadata(ctx context.Context, req *ProviderPackageMetadataRequest) (*ProviderPackageMetadataResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
