package util

import "os"

func GetVersion() string {
	data, err := os.ReadFile("./version.txt")
	if err != nil {
		return ""
	}

	return string(data)
}
