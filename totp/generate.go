package totp

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

// Generates a random secret key that is used to generate the TOTP (gets sent to the user)
func GenerateSecret() (string, error) {
	var secret [10]byte
	_, err := rand.Read(secret[:])
	if err != nil {
		return "", err
	}
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(secret[:]), nil
}

// Generates a TOTP using the secret key and the current time (basically what the user would do)
func GenerateTOTP(secret string) string {
	counter := time.Now().Unix() / 30
	hash := hmacSHA1(secret, counter)
	otp := truncate(hash)
	return fmt.Sprintf("%06d", otp)
}

func hmacSHA1(secret string, counter int64) []byte {
	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		panic(err)
	}
	counterBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(counterBytes, uint64(counter))
	hash := hmac.New(sha1.New, key)
	hash.Write(counterBytes)
	return hash.Sum(nil)
}

func truncate(hash []byte) int {
	offset := hash[len(hash)-1] & 0x0f
	truncatedHash := hash[offset : offset+4]
	truncatedHash[0] = truncatedHash[0] & 0x7f
	return int(binary.BigEndian.Uint32(truncatedHash)) % 1000000
}
