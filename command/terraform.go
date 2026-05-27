package command

// TerraformCommand is a command which update version constraints for terraform.
type TerraformCommand struct {
	Meta
	version     string
	path        string
	recursive   bool
	ignorePaths []string
}

// Run runs the procedure of this command.
func (c *TerraformCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// Help returns long-form help text.
func (c *TerraformCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *TerraformCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
