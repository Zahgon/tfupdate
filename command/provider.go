package command

// ProviderCommand is a command which update version constraints for provider.
type ProviderCommand struct {
	Meta
	name        string
	version     string
	path        string
	recursive   bool
	ignorePaths []string
}

// Run runs the procedure of this command.
func (c *ProviderCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Help returns long-form help text.
func (c *ProviderCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *ProviderCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
