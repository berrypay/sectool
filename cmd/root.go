/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /cmd/root.go
 * Created Date: Friday September 1st 2023 11:06:43 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday September 1st 2023 15:24:13 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package cmd

import (
	"os"

	"github.com/berrypay/sectool/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sectool",
	Short: "Utilities for application security setup and usage",
	Long: `BerryPay sectool provide set of tools for setting up application security components
such as encrypting application secrets, generating application keys and certificates, etc.`,
	Version: app.AppVersion,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sectool.yaml)")
	rootCmd.PersistentFlags().BoolP(app.DEBUG_FLAG, "d", false, "Enable debug mode")
	viper.BindPFlag(app.DEBUG_FLAG, rootCmd.PersistentFlags().Lookup(app.DEBUG_FLAG))
	rootCmd.PersistentFlags().StringP(app.PRIVATE_KEY_FLAG, "k", "app.key", "Application Key (PEM) file")
	viper.BindPFlag(app.PRIVATE_KEY_FLAG, rootCmd.PersistentFlags().Lookup(app.PRIVATE_KEY_FLAG))
	rootCmd.PersistentFlags().StringP(app.CERT_FLAG, "c", "app.crt", "Application Certificate (PEM) file")
	viper.BindPFlag(app.CERT_FLAG, rootCmd.PersistentFlags().Lookup(app.CERT_FLAG))
	rootCmd.PersistentFlags().StringP(app.OAEP_LABEL_FLAG, "l", "Our Application Secret", "OAEP Label")
	viper.BindPFlag(app.OAEP_LABEL_FLAG, rootCmd.PersistentFlags().Lookup(app.OAEP_LABEL_FLAG))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
