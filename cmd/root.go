// Copyright © 2018 Ercan Aydogan
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

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/eaydogan/palto/cli"
)

var cfgFile string
var opt cli.ScanOption
// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "palto",
	Short: "Palto, is a Kubernetes public faces services security checker.",
	Long: `Palto, search kubernetes known services ports and try to get
information from that services. For example:

etcd is a distributed key-value store. Like mongoDB default installation is open public access.
With Palto, you can search your etcd status. This application is a tool to generate the needed files (kubeconfig, token secret)
to quickly check etc services.`,

		Run: func(cmd *cobra.Command, args []string) { 
            cli.Scan(&opt)
		},
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
	opt.StartPort=2379
	opt.StopPort=2379
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.palto.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.PersistentFlags().StringVar(&opt.IP, "ip", "", "search ip address.")
	
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

		// Search config in home directory with name ".palto" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".palto")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
