package main

import (
	"fmt"
	"totp/totp"
)

func main() {
	secret, err := totp.GenerateSecret()
	if err != nil {
		fmt.Println("Error generating secret:", err)
		return
	}
	fmt.Println("Secret:", secret)

	currentTOTP := totp.GenerateTOTP(secret)
	fmt.Println("Current TOTP:", currentTOTP)

	for {
		fmt.Print("Enter the TOTP: ")
		var userTOTP string
		fmt.Scanln(&userTOTP)

		// Validate the provided TOTP
		if totp.ValidateTOTP(secret, userTOTP) {
			fmt.Println("🟢 TOTP is valid!")
		} else {
			fmt.Println("🔴 TOTP is invalid!")
		}
	}
}
