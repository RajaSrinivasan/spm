/*
Copyright © 2020 rs@toprllc.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var Pubpkg string
var Pubart string
var PublicKey, PrivateKey, PkiPassphrase string
var PkgFormat, PkgPassword string
var Workarea string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "spm",
	Short: "Secure (simple) package manager",
	Long: `Secure package manager helps prepare and distribute packages of applications 
and/or data.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spm.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Home dir is %s\n", home)
		// Search config in home directory with name ".spm" (without extension).
		//viper.AddConfigPath(home)
		viper.AddConfigPath("./example")
		viper.SetConfigName(".spm")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		Pubpkg = viper.GetString("pubpkg")
		Pubart = viper.GetString("pubart")
		fmt.Printf("Pkg publish url=%s Artifacts=%s\n", Pubpkg, Pubart)
		PublicKey = viper.GetString("security.publickey")
		PrivateKey = viper.GetString("security.privatekey")
		PkgFormat = viper.GetString("package.format")
		viper.SetEnvPrefix("spm")
		viper.BindEnv("pkipassphrase")
		viper.BindEnv("pkgpassword")
		PkgPassword = viper.GetString("pkgpassword")
		PkiPassphrase = viper.GetString("pkipassphrase")
		Workarea = viper.GetString("package.workarea")
		fmt.Printf("PKI passphrase %s Pkg Password %s Workarea %s\n", PkiPassphrase, PkgPassword, Workarea)

	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}
