/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /app/security/encrypt-aes.go
 * Created Date: Tuesday September 5th 2023 12:17:56 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 12:52:36 +0800
 * Modified By: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * --------------------------------------
 * Copyright (c) 2023 BerryPay (M) Sdn. Bhd.
 */

package security

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/berrypay/appsec"
)

func EncryptAES(secret string, key string) {
	keyBytes := []byte(key)
	err := appsec.InitAES(keyBytes)
	if err != nil {
		fmt.Printf("Error initiating AES-GCM: %s\n", err.Error())
		os.Exit(5)
	}

	cipherText, err := appsec.EncryptAESGCM([]byte(secret))
	if err != nil {
		fmt.Printf("Error encrypting secret: %s\n", err.Error())
		os.Exit(6)
	}

	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(cipherText))
}
