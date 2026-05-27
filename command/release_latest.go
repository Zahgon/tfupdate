package command

// ReleaseLatestCommand is a command which gets the latest release version.
type ReleaseLatestCommand struct {
	Meta
	sourceType string
	source     string
}

// Run runs the procedure of this command.
func (c *ReleaseLatestCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Help returns long-form help text.
func (c *ReleaseLatestCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *ReleaseLatestCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
