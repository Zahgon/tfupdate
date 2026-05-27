package command

// ReleaseListCommand is a command which gets a list of release versions.
type ReleaseListCommand struct {
	Meta
	maxLength  int
	preRelease bool
	sourceType string
	source     string
}

// Run runs the procedure of this command.
func (c *ReleaseListCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Help returns long-form help text.
func (c *ReleaseListCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *ReleaseListCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
