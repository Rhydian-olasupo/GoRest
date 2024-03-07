package main

import (
	"go_trial/gorest/encrypString/utils"
	"log"
)

func main() {
	key := "111023043350789514532147"
	message := "I am a message"

	log.Println("Original Message: ", message)
	encryptedString := utils.EncryptString(key, message)
	log.Println("Encrypted message: ", encryptedString)
	decryptedString := utils.DecryptString(key, encryptedString)
	log.Println("Decrypted message: ", decryptedString)
}
