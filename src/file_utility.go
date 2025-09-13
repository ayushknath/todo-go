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

func readFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error occured while reading JSON file")
		os.Exit(1)
	}
	return data
}
