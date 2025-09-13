package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func jsonEncode(data *[]TodoItem) []byte {
	jsonData, err := json.Marshal(*data)
	if err != nil {
		fmt.Println("Error occured while encoding JSON data")
		os.Exit(1)
	}
	return jsonData
}

func jsonDecode(jsonData []byte, data *[]TodoItem) {
	err := json.Unmarshal(jsonData, data)
	if err != nil {
		fmt.Println("Error occured while decoding JSON data")
		os.Exit(1)
	}
}
