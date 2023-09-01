/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /cmd/encrypt.go
 * Created Date: Friday September 1st 2023 11:44:30 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday September 1st 2023 15:43:14 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package cmd

import (
	"github.com/berrypay/sectool/app"
	"github.com/berrypay/sectool/app/security"
	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt secret_string",
	Short: "Encrypt a secret",
	Long: `Encrypt a given string using the specified public key with RSA-OAEP.
By default it use "Our Application Secret" as the OAEP label.

Returns a base64 encoded string of the secret cipher.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		security.Encrypt(args[0])
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// encryptCmd.Flags().StringP("secret", "s", "", "Secret to encrypt")
	// encryptCmd.MarkFlagRequired("secret")
	encryptCmd.MarkFlagRequired(app.OAEP_LABEL_FLAG)
}
