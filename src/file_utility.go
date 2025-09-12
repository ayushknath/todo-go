package main

import (
	"fmt"
	"os"
)

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		} else {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
	return true
}
