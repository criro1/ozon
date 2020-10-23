package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	cost = "cost"
	quantity = "quantity"
)
func main() {
	file, err := os.Open("data.csv")
	// если запускаете из GoLand, возможно будет ошибка "no such file or directory"
	// тогда заменить на строку ниже, или изменить путь на абсоютный
	//file, err := os.Open("github.com/criro1/ozon/data.csv")
	defer file.Close()
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.Exit(0)
	}

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 6

	mp := make(map[string]map[string]int)

	var (
		tmp int
	)
	for {
		record, err := reader.Read()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		if record[0] == "id" {
			continue
		}
		if _, ok := mp[record[2]]; !ok {
			mp[record[2]] = make(map[string]int)
		}

		tmp, err = strconv.Atoi(record[3])
		if err != nil {
			continue
		}
		mp[record[2]][cost] += tmp

		tmp, err = strconv.Atoi(record[4])
		if err != nil {
			continue
		}
		mp[record[2]][quantity] += tmp
	}

	for key, value := range mp {
		fmt.Println("Good name - '" + key + "':", "cost =", value[cost], "; quantity =", value[quantity])
	}
}