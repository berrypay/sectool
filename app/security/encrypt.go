/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /app/security/encrypt.go
 * Created Date: Friday September 1st 2023 12:15:54 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Friday September 1st 2023 15:52:36 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package security

import (
	"fmt"

	"github.com/berrypay/appsec"
	"github.com/berrypay/sectool/app"
	"github.com/spf13/viper"
)

func Encrypt(secret string) {
	if viper.GetBool(app.DEBUG_FLAG) {
		fmt.Println("Encrypting secret...")
		fmt.Printf("Using OAEP label: %s\n", viper.GetString(app.OAEP_LABEL_FLAG))
		fmt.Printf("Given secret: %s\n", secret)
	}

	err := appsec.LoadPublicKey(viper.GetString(app.CERT_FLAG))
	if err != nil {
		fmt.Printf("Error loading public key: %s\n", err.Error())
	}

	cipherText, err := appsec.EncryptOAEP(secret, viper.GetString(app.OAEP_LABEL_FLAG))
	if err != nil {
		fmt.Printf("Error encrypting secret: %s\n", err.Error())
	}

	if viper.GetBool(app.DEBUG_FLAG) {
		fmt.Print("Encrypted secret: ")
	}
	fmt.Printf("%s\n", cipherText)
}
