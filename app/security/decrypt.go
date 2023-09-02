/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /app/security/decrypt.go
 * Created Date: Friday September 1st 2023 12:17:19 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Saturday September 2nd 2023 10:37:20 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package security

import (
	"crypto/rsa"
	"fmt"
	"os"

	"github.com/berrypay/appsec"
	"github.com/berrypay/sectool/app"
	"github.com/spf13/viper"
)

var AppPublicKey *rsa.PublicKey

func Decrypt(cipher string) {
	if viper.GetBool(app.DEBUG_FLAG) {
		fmt.Println("Decrypting secret...")
		fmt.Printf("Using OAEP label: %s\n", viper.GetString(app.OAEP_LABEL_FLAG))
		fmt.Printf("Given cipher: %s\n", cipher)
	}

	err := appsec.LoadPrivateKey(viper.GetString(app.PRIVATE_KEY_FLAG))
	if err != nil {
		fmt.Printf("Error loading private key: %s\n", err.Error())
		os.Exit(3)
	}

	plainText, err := appsec.DecryptOAEP(cipher, viper.GetString(app.OAEP_LABEL_FLAG))
	if err != nil {
		fmt.Printf("Error decrypting cipher: %s\n", err.Error())
		os.Exit(4)
	}

	if viper.GetBool(app.DEBUG_FLAG) {
		fmt.Print("Decrypted secret: ")
	}
	fmt.Printf("%s\n", plainText)
}
