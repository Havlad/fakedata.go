package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-faker/faker/v4"
)

type User struct {
	Name      string `faker:"first_name"`
	Familia   string `faker:"last_name"`
	Otchestvo string `faker:"name"`
	Birthday  string `faker:"date"`
	Age       string `faker:"year"`
}

func main() {
	user := User{}
	err := faker.FakeData(&user)
	if err != nil {
		fmt.Println("Error faker", err)
		return
	}
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	err = os.WriteFile("user.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("JSON data written to file successfully.")
}
