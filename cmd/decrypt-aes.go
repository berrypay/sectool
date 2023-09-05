/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /cmd/decrypt-aes.go
 * Created Date: Tuesday September 5th 2023 12:34:15 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 12:54:10 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package cmd

import (
	"github.com/berrypay/sectool/app/security"
	"github.com/spf13/cobra"
)

// decryptAesCmd represents the decryptAes command
var decryptAesCmd = &cobra.Command{
	Use:   "decryptAes",
	Short: "Decrypt a base64 encoded cipher text using AES-GCM",
	Long: `Decrypt a given base64 encoded cipher text using the specified password as key with AES-GCM method.
	
Returns the plain text.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cmd.Flags().GetString("password")
		security.DecryptAES(args[0], password)
	},
}

func init() {
	rootCmd.AddCommand(decryptAesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptAesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptAesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	decryptAesCmd.Flags().StringP("password", "p", "", "32 bytes AES Secret Key")
	decryptAesCmd.MarkFlagRequired("password")
}
