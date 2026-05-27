package release

import (
	version "github.com/hashicorp/go-version"
)

func tagNameToVersion(tagName string) string {
	_ = "STUB: not implemented"
	// if a tagName starts with `v`, remove it.
	return ""
}

func reverseStringSlice(s []string) []string {
	_ = "STUB: not implemented"

	// apparently inefficient but simple way
	return nil
}

func minInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

// toVersions converts []string to []*version.Version.
// Ignore if parse error.
func toVersions(versionsRaw []string) []*version.Version { _ = "STUB: not implemented"; return nil }

// fromVersions converts []*version.Version to []string.
func fromVersions(versions []*version.Version) []string { _ = "STUB: not implemented"; return nil }

// sortVersions sort a list of versions in semver order.
func sortVersions(versions []*version.Version) []*version.Version {
	_ = "STUB: not implemented"
	return nil
}

// excludePreReleases excludes pre-releases such as alpha, beta, rc.
func excludePreReleases(versions []*version.Version) []*version.Version {
	_ = "STUB: not implemented"
	// exclude pre-release
	return nil
}
