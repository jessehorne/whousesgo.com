package util

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func GenerateAdminToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

// UpdateAdminToken generates a random cryptographically secure token at a certain interval and
// makes it accessible as an environment variable.
func StartUpdateAdminTokenJob(howOften time.Duration) {
	ticker := time.NewTicker(howOften)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				newAdminToken, err := GenerateAdminToken()
				if err != nil {
					newAdminToken = "NOALLOW"
				}
				if err := os.Setenv("ADMIN_TOKEN", newAdminToken); err != nil {
					fmt.Println("couldn't set ADMIN_TOKEN", err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
