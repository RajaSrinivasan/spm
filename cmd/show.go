package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the contantes of the package",
	Long:  `Unpack the contents of the package, verify and list details about the package`,
	Run:   show,
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func show(cmd *cobra.Command, args []string) {
	fmt.Println("show called")
}
