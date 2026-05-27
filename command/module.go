package command

// ModuleCommand is a command which update version constraints for module.
type ModuleCommand struct {
	Meta
	name            string
	version         string
	path            string
	recursive       bool
	ignorePaths     []string
	sourceMatchType string
}

// Run runs the procedure of this command.
func (c *ModuleCommand) Run(args []string) int { _ = "STUB: not implemented"; return 0 }

// For modules, automatic latest version resolution is not simple.
// To implement, we will probably need to get information from the Terraform Registry.

// Help returns long-form help text.
func (c *ModuleCommand) Help() string { _ = "STUB: not implemented"; return "" }

// Synopsis returns one-line help text.
func (c *ModuleCommand) Synopsis() string { _ = "STUB: not implemented"; return "" }
