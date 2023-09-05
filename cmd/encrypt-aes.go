/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /cmd/encrypt-aes.go
 * Created Date: Tuesday September 5th 2023 12:12:24 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 12:46:50 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package cmd

import (
	"github.com/berrypay/sectool/app/security"
	"github.com/spf13/cobra"
)

// encryptAesCmd represents the encryptAes command
var encryptAesCmd = &cobra.Command{
	Use:   "encryptAes",
	Short: "Encrypt a secret using AES-GCM",
	Long: `Encrypt a given string using the specified password as key for using AES-GCM method.

Returns a base64 encoded string of the secret cipher.`,
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cmd.Flags().GetString("password")
		security.EncryptAES(args[0], password)
	},
}

func init() {
	rootCmd.AddCommand(encryptAesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptAesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptAesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	encryptAesCmd.Flags().StringP("password", "p", "", "32 bytes AES Secret Key")
	encryptAesCmd.MarkFlagRequired("password")
}
