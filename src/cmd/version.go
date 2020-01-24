package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Report version",
	Long:  `Report version`,
	Run:   version,
}

func init() {

	rootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	fmt.Printf("Version %d.%d-%d Built %s\n", versionMajor, versionMinor, versionBuild, buildTime)
}
