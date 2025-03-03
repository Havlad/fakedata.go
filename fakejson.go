package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		log.Fatalf("Использование: go run main.go <количество> <файл>")
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count <= 0 {
		log.Fatalf("Ошибка: введите корректное число пользователей")
	}

	filename := os.Args[2]
	users := make([]User, count)

	for range count {
		err = faker.FakeData(&users)
		if err != nil {
			log.Fatalf("Проблема при факировании данных")
		}
	}

	data, _ := json.Marshal(users)
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Ошибка записи файла")
	}

	fmt.Println("Файл", filename, "успешно создан c", count, `пользователями!`)
}
