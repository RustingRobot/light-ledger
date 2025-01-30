package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Expenses expenses `json:"expenses"`
}

type expenses struct {
	Description []string `json:"desc"`
	Cost        []string `json:"cost"`
}

func SaveToFile(data Data) {
	byte_data, _ := json.Marshal(data)
	if err := os.WriteFile("db.json", []byte(byte_data), 0666); err != nil {
		fmt.Println("couldn't save to file")
	}
}
