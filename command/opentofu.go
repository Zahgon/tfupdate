package command

// OpenTofuCommand is a command which update version constraints for OpenTofu.
type OpenTofuCommand struct {
	Meta
	version     string
	path        string
	recursive   bool
	ignorePaths []string
}

// Run runs the procedure of this command.
func (c *OpenTofuCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Help returns long-form help text.
func (c *OpenTofuCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *OpenTofuCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
