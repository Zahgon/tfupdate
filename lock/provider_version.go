package lock

// ProviderVersion is a data structure that holds hash values of a specific
// version of a particular provider. It corresponds to one provider block in
// the dependency lock file (.terraform.lock.hcl).
// https://developer.hashicorp.com/terraform/language/files/dependency-lock
type ProviderVersion struct {
	// address is a provider address such as hashicorp/null.
	address string

	// version is a version number such as 3.2.1.
	version string

	// platforms is a list of target platforms to generate hash values.
	// Target platform names consist of an operating system and a CPU architecture such as darwin_arm64.
	// The actual lock file does not distinguish which platform the hash values
	// belong to, but we keep them distinct in memory for easy debugging in case
	// of checksum mismatches.
	platforms []string

	// h1Hashes is a dictionary of hash values calculated with the h1 scheme.
	// The key is the filename.
	h1Hashes map[string]string

	// zhHashes is a dictionary of hash values calculated with the zh scheme.
	// The key is the filename.
	zhHashes map[string]string
}

// newEmptyProviderVersion returns a new empty ProviderVersion, which is
// intended to be used as a variable to store merge results.
func newEmptyProviderVersion(address string, version string) *ProviderVersion {
	_ = "STUB: not implemented"
	return nil
}

// Merge takes another ProviderVersion and merges it. It returns an error if
// the argument is incompatible the current object.
func (pv *ProviderVersion) Merge(rhs *ProviderVersion) error { _ = "STUB: not implemented"; return nil }

// should not happen

// AllHashes returns an array of strings containing all hash values. It is
// intended to be used as the value of hashes in a dependency lock file.
// The result is sorted alphabetically.
func (pv *ProviderVersion) AllHashes() []string { _ = "STUB: not implemented"; return nil }
