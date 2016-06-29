package commands

import (
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processor"
	"github.com/nanobox-io/nanobox/util/print"
	"github.com/nanobox-io/nanobox/validate"
)

var (

	// DestroyCmd ...
	DestroyCmd = &cobra.Command{
		Use:   "destroy",
		Short: "Start the virtual machine",
		Long: `
Start the virtual machine.
		`,
		PreRun: validate.Requires("provider"),
		Run:    destroyFn,
	}
)

// destroyFn ...
func destroyFn(ccmd *cobra.Command, args []string) {
	print.OutputCommandErr(processor.Run("destroy", processor.DefaultControl))
}
