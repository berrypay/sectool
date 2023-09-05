/*
 * Project: BerryPay Application Security Tool Set
 * Filename: /app/security/decrypt-aes.go
 * Created Date: Tuesday September 5th 2023 12:18:21 +0800
 * Author: Sallehuddin Abdul Latif (sallehuddin@berrypay.com)
 * Company: BerryPay (M) Sdn. Bhd.
 * --------------------------------------
 * Last Modified: Tuesday September 5th 2023 12:56:12 +0800
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

func DecryptAES(cipher string, key string) {
	keyBytes := []byte(key)
	err := appsec.InitAES(keyBytes)
	if err != nil {
		fmt.Printf("Error initiating AES-GCM: %s\n", err.Error())
		os.Exit(5)
	}

	decodedCipher, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		fmt.Printf("Error decoding cipher text: %s\n", err.Error())
		os.Exit(7)
	}

	plainText, err := appsec.DecryptAESGCM([]byte(decodedCipher))
	if err != nil {
		fmt.Printf("Error decrypting secret: %s\n", err.Error())
		os.Exit(8)
	}

	fmt.Printf("%s\n", plainText)
}
