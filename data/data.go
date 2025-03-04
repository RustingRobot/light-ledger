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

// sort.Interface sorts by date
func (d Data) Len() int      { return len(d.Expenses) }
func (d Data) Swap(i, j int) { d.Expenses[i], d.Expenses[j] = d.Expenses[j], d.Expenses[i] }
func (d Data) Less(i, j int) bool {
	return d.Expenses[i].Date < d.Expenses[j].Date
}

func SaveToFile(data Data, tags []string, adding bool) {
	for _, e := range tags {
		if adding {
			data.Tags[e] += 1
		} else {
			data.Tags[e] -= 1
			if data.Tags[e] < 1 {
				delete(data.Tags, e)
			}
		}
	}
	byte_data, _ := json.Marshal(data)
	if err := os.WriteFile("db.json", []byte(byte_data), 0666); err != nil {
		fmt.Println("couldn't save to file")
	}
}
