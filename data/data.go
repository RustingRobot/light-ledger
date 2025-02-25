package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Expenses []Entry        `json:"expenses"`
	Tags     map[string]int `json:"tags"`
}

type Entry struct {
	Description string   `json:"desc"`
	Cost        string   `json:"cost"`
	Date        string   `json:"date"`
	Tags        []string `json:"tags"`
}

func SaveToFile(data Data, tags []string, adding bool) {
	fmt.Println(tags)
	for _, e := range tags {
		fmt.Println(e)
		if adding {
			data.Tags[e] += 1
		} else {
			data.Tags[e] -= 1
			if data.Tags[e] < 1 {
				delete(data.Tags, e)
			}
		}
	}
	fmt.Println(data)
	byte_data, _ := json.Marshal(data)
	if err := os.WriteFile("db.json", []byte(byte_data), 0666); err != nil {
		fmt.Println("couldn't save to file")
	}
}
