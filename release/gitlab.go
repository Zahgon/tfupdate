package release

import (
	"context"

	"github.com/xanzy/go-gitlab"
)

// GitLabAPI is an interface which calls GitLab API.
// This abstraction layer is needed for testing with mock.
type GitLabAPI interface {
	// ProjectListReleases gets a pagenated of releases accessible by the authenticated user.
	ProjectListReleases(ctx context.Context, owner, project string, opt *gitlab.ListReleasesOptions) ([]*gitlab.Release, *gitlab.Response, error)
}

// GitLabConfig is a set of configurations for GitLabRelease..
type GitLabConfig struct {
	// api is an instance of GitLabAPI interface.
	// It can be replaced for testing.
	api GitLabAPI

	// BaseURL is a URL for GitLab API requests.
	// Defaults to the public GitLab API.
	// BaseURL should always be specified with a trailing slash.
	BaseURL string

	// Token is a personal access token for GitLab, needed to use the api.
	Token string
}

// GitLabClient is a real GitLabAPI implementation.
type GitLabClient struct {
	client *gitlab.Client
}

var _ GitLabAPI = (*GitLabClient)(nil)

// NewGitLabClient returns a real GitLab instance.
func NewGitLabClient(config GitLabConfig) (*GitLabClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ProjectListReleases gets a pagenated of releases accessible by the authenticated user.
func (c *GitLabClient) ProjectListReleases(ctx context.Context, owner, project string, opt *gitlab.ListReleasesOptions) ([]*gitlab.Release, *gitlab.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GitLabRelease is a release implementation which provides version information with GitLab Release.
type GitLabRelease struct {
	// api is an instance of GitLabAPI interface.
	// It can be replaced for testing.
	api GitLabAPI

	// owner is a namespace of project.
	// limited to one level (group or personal - not sub-groups?)
	owner string

	// project is a name of project (repository).
	project string
}

var _ Release = (*GitLabRelease)(nil)

// NewGitLabRelease is a factory method which returns an GitLabRelease instance.
func NewGitLabRelease(source string, config GitLabConfig) (*GitLabRelease, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// If config.api is not set, create a default GitLabClient

// ListReleases returns a list of unsorted all releases including pre-release.
func (r *GitLabRelease) ListReleases(ctx context.Context) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// max
