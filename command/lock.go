package command

// LockCommand is a command which update dependency lock files.
type LockCommand struct {
	Meta
	platforms   []string
	path        string
	recursive   bool
	ignorePaths []string
}

// Run runs the procedure of this command.
func (c *LockCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Fetch environment variables

// Create tfregistry.Config

// Help returns long-form help text.
func (c *LockCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *LockCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
