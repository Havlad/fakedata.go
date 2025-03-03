package main

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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
	if len(os.Args) < 3 {
		fmt.Println("Использование: go run main.go <количество> <файл>")
		return
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count <= 0 {
		fmt.Println("Ошибка: введите корректное число пользователей")
		return
	}

	filename := os.Args[2]
	users := make([]User, count)

	for i := 0; i < count; i++ {
		_ = faker.FakeData(&users[i])
	}

	data, _ := json.MarshalIndent(users, "", "  ")
	_ = os.WriteFile(filename, data, 0644)

	fmt.Println("Файл", filename, "успешно создан с", count, "пользователями!")
}
