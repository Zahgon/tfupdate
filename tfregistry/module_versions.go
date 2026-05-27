package tfregistry

import (
	"context"
)

// ListModuleVersionsRequest is a request parameter for the ListModuleVersions API.
type ListModuleVersionsRequest struct {
	// The user or organization the module is owned by.
	Namespace string `json:"namespace"`
	// The name of the module.
	Name string `json:"name"`
	// The name of the provider.
	Provider string `json:"provider"`
}

// ListModuleVersionsResponse is a response data for the ListModuleVersions API.
type ListModuleVersionsResponse struct {
	// Modules is an array containing module information.
	// The first element contains the requested module.
	Modules []ModuleVersions `json:"modules"`
}

// ModuleVersions represents version information for a module.
type ModuleVersions struct {
	// Versions is a list of available versions.
	Versions []ModuleVersion `json:"versions"`
}

// ModuleVersion represents a single version of a module.
type ModuleVersion struct {
	// Version is the version string.
	Version string `json:"version"`
}

// ListModuleVersions returns all versions of a module for a single provider.
// This works for both Terraform and OpenTofu registries.
func (c *Client) ListModuleVersions(ctx context.Context, req *ListModuleVersionsRequest) (*ListModuleVersionsResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
