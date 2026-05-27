package release

import (
	"context"
	"net/http"

	"github.com/google/go-github/v28/github"
)

// GitHubAPI is an interface which calls GitHub API.
// This abstraction layer is needed for testing with mock.
type GitHubAPI interface {
	// RepositoriesListReleases lists the releases for a repository.
	// GitHub API docs: https://developer.github.com/v3/repos/releases/#list-releases-for-a-repository
	RepositoriesListReleases(ctx context.Context, owner, repo string, opt *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error)
}

// GitHubConfig is a set of configurations for GitHubRelease.
type GitHubConfig struct {
	// api is an instance of GitHubAPI interface.
	// It can be replaced for testing.
	api GitHubAPI

	// BaseURL is a URL for GitHub API requests.
	// Defaults to the public GitHub API.
	// This looks like the GitHub Enterprise support, but currently for testing purposes only.
	// The GitHub Enterprise is not supported yet.
	// BaseURL should always be specified with a trailing slash.
	BaseURL string

	// Token is a personal access token for GitHub.
	// This allows access to a private repository.
	Token string
}

// GitHubClient is a real GitHubAPI implementation.
type GitHubClient struct {
	client *github.Client
}

var _ GitHubAPI = (*GitHubClient)(nil)

// NewGitHubClient returns a real GitHubClient instance.
func NewGitHubClient(config GitHubConfig) (*GitHubClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// newOAuth2Client returns a *http.Client which sets a given token to the Authorization header.
// This allows access to a private repository.
func newOAuth2Client(token string) *http.Client { _ = "STUB: not implemented"; return nil }

// RepositoriesListReleases lists the releases for a repository.
func (c *GitHubClient) RepositoriesListReleases(ctx context.Context, owner, repo string, opt *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GitHubRelease is a release implementation which provides version information with GitHub Release.
type GitHubRelease struct {
	// api is an instance of GitHubAPI interface.
	// It can be replaced for testing.
	api GitHubAPI

	// owner is a namespace of repository.
	owner string

	// repo is a name of repository.
	repo string
}

var _ Release = (*GitHubRelease)(nil)

// NewGitHubRelease is a factory method which returns an GitHubRelease instance.
func NewGitHubRelease(source string, config GitHubConfig) (Release, error) {
	_ = "STUB: not implemented"
	return *new(Release), nil
}

// If config.api is not set, create a default GitHubClient

// ListReleases returns a list of unsorted all releases including pre-release.
func (r *GitHubRelease) ListReleases(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// max
