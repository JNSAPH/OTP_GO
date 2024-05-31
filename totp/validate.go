package totp

func ValidateTOTP(secret, userTOTP string) bool {
	generatedTOTP := GenerateTOTP(secret)
	return userTOTP == generatedTOTP
}
