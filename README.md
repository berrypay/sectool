# Sectool: BerryPay Application Security Tool Set

Sectool is a set of tools for setting up application security components such as encrypting application secrets, generating application keys and certificates, etc.

## Build

1. Clone the repository

   ```console
   foo@bar:~$ git clone https://github.com/berrypay/sectool.git
   ```

2. Install all dependency

   ```console
   foo@bar:~$ cd sectool
   foo@bar:~$ go get ./...
   ```

3. Build binary

   ```console
   foo@bar:~$ make
   ```

4. Use the tool

   ```console
   foo@bar:~$ ./sectool -h
   BerryPay sectool provide set of tools for setting up application security components
   such as encrypting application secrets, generating application keys and certificates, etc.

   Usage:
   sectool [command]

   Available Commands:
   completion  Generate the autocompletion script for the specified shell
   decrypt     Decrypt a cipher text
   decryptAes  Decrypt a base64 encoded cipher text using AES-GCM
   encrypt     Encrypt a secret
   encryptAes  Encrypt a secret using AES-GCM
   help        Help about any command
   version     Show version of sectool

   Flags:
   -d, --debug                Enable debug mode
   -h, --help                 help for sectool
   -l, --oaep-label string    OAEP Label (default "Our Application Secret")
   -k, --private-key string   Application Key (PEM) file (default "app.key")
   -c, --public-cert string   Application Certificate (PEM) file (default "app.crt")
   -v, --version              version for sectool

   Use "sectool [command] --help" for more information about a command.
   ```
