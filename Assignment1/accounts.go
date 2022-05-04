//this program takes a json file name in command line arguments and update the balance in each record
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Account struct {
	Name      string
	AccountID int
	Balance   float64
}

func main() {
	var result []Account
	data, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error in opening file!")
	}
	d := json.NewDecoder(data)
	d.Decode(&result)
	for i := 0; i < len(result); i++ {
		result[i].Balance = result[i].Balance + 100
	}
	file, _ := json.MarshalIndent(result, "", " ")
	ioutil.WriteFile("accountsUpdated.json", file, 0644)

	data.Close()
}
