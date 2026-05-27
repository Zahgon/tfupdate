package lock

import (
	"os"
)

// zipDataToH1Hash is a helper function that calculates the h1 hash value from
// bytes sequence of the provider's zip archive.
func zipDataToH1Hash(zipData []byte) (string, error) { _ = "STUB: not implemented"; return "", nil }

// The h1 hash value in .terraform.lock.hcl uses the same hash function as go.sum.

// writeTempFile writes content to a temporary file and return its file.
func writeTempFile(content []byte) (*os.File, error) { _ = "STUB: not implemented"; return nil, nil }

// shaSumsDataToZhHash is a helper function for parsing zh hash values from
// bytes sequence of the shaSumsData document.
func shaSumsDataToZhHash(shaSumsData []byte) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read an entry per line.

// We expect that blank lines are not normally included, but to make the
// test data easier to read, ignore blank lines.

// Split rows into columns with spaces, but note that there are two spaces between the columns.
// e4453fbebf90c53ca3323a92e7ca0f9961427d2f0ce0d2b65523cc04d5d999c2  terraform-provider-null_3.2.1_darwin_arm64.zip

// Initially, we thought of using the key of the zh hash as the platform,
// but we found out that it also includes metadata such as manifest.json,
// so we decided to use the filename as it is.

// As the implementation of the h1 hash includes a prefix for the "h1:"
// scheme, zh also includes the "zh:" prefix for consistency.
