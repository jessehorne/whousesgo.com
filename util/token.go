package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func IsValidToken(token string) bool {
	data, err := os.ReadFile("./token.txt")
	if err != nil {
		return false
	}

	if string(data) != token {
		return false
	}

	return true
}

func GenerateAdminToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func WriteToken() {
	newAdminToken, err := GenerateAdminToken()
	if err != nil {
		newAdminToken = "NOALLOW"
	}

	data := []byte(newAdminToken)
	if err := os.WriteFile("./token.txt", data, 0644); err != nil {
		fmt.Println("ERROR: couldn't set ADMIN_TOKEN", err)
	}

	fmt.Println(IsValidToken(newAdminToken))
}

// UpdateAdminToken generates a random cryptographically secure token at a certain interval and
// makes it accessible as an environment variable.
func StartUpdateAdminTokenJob(howOften time.Duration) {
	WriteToken()

	ticker := time.NewTicker(howOften)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				WriteToken()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
