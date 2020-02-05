package cmd

import (
	"github.com/RajaSrinivasan/spm/impl"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a secure package",
	Long: `Create a secure package based on the configuration file provided.
Optionally push the artifact(s) to a distribution server. 
The first argument is the package spec file (ex spec.yaml)
Output package name is the second argument`,
	Args: cobra.MinimumNArgs(2),
	Run:  build,
}

func init() {
	rootCmd.AddCommand(buildCmd)
}

func build(cmd *cobra.Command, args []string) {
	impl.Build(args[0], args[1])
}
