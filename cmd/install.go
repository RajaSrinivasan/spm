package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install the package",
	Long: `Install the package provided first verifying the integrity of the artifacts. Argument
	is the package (.spm)%s\n`,
	Args: cobra.MinimumNArgs(1),
	Run:  install,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) {
	fmt.Printf("install called %s\n", args[0])
}
