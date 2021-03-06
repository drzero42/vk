// Copyright © 2019 Cellpoint Mobile
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	glogcobra "github.com/blocktop/go-glog-cobra"
	"github.com/cellpointmobile/vk/program"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	force   bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vk",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.vk/config.yaml)")
	rootCmd.PersistentFlags().StringP("bindir", "b", "$HOME/.local/bin", "Directory for bin-files.")
	rootCmd.PersistentFlags().String("definitions", "", "URL/path to definitions file.")
	rootCmd.PersistentFlags().BoolVar(&program.ClearCache, "clear-cache", false, "clear the cache.")

	viper.BindPFlag("bindir", rootCmd.PersistentFlags().Lookup("bindir"))
	viper.SetDefault("bindir", "$HOME/.local/bin")
	viper.BindPFlag("definitions", rootCmd.PersistentFlags().Lookup("definitions"))
	viper.SetDefault("definitions", "https://raw.githubusercontent.com/cellpointmobile/vk-definitions/master/vk-definitions.json")

	glogcobra.Init(rootCmd)
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

		// Search config in home directory with name ".vk" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".vk/config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.ReadInConfig()
	// If a config file is found, read it in.
	//if err := viper.ReadInConfig(); err == nil {
	//	glog.Infof("Using config file: %s\n", viper.ConfigFileUsed())
	//}
	glogcobra.Parse(rootCmd)
}
