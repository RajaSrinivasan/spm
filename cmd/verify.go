package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify the package",
	Long: `Unpack the package and authenticate the contents against the public key in the package.
	Argument is the package (.spm).
	The package will not be installed. The unpacked contents will be left in the work area for inspection.\n`,
	Args: cobra.MinimumNArgs(1),
	Run:  install,
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}

func verify(cmd *cobra.Command, args []string) {
	fmt.Printf("verify called %s\n", args[0])
}
