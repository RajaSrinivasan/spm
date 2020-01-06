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
Optionally push the artifact(s) to a distribution server`,
	Run: build,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func build(cmd *cobra.Command, args []string) {
	fmt.Println("build called")
}
