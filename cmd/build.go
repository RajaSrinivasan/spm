package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a secure package",
	Long: `Create a secure package based on the configuration file provided.
Optionally push the artifact(s) to a distribution server. The argument is the package spec file (ex spec.yaml)`,
	Args: cobra.MinimumNArgs(1),
	Run:  build,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func build(cmd *cobra.Command, args []string) {
	fmt.Printf("build called to process %s\n", args[0])
}
