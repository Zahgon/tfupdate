package command

// ReleaseCommand is a command which just shows help for subcommands.
type ReleaseCommand struct {
	Meta
}

// Run runs the procedure of this command.
func (c *ReleaseCommand) Run(args []string) int {
	_ = "STUB: not implemented" // nolint revive unused-parameter
	return 0
}

// Help returns long-form help text.
func (c *ReleaseCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *ReleaseCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
