package main

import (
	database "InterfaceDroch/interface/go"
	"context"
	"fmt"
)

func main() {

	db := database.NewData()

	ctx := context.Background()

	err := db.Set(ctx, "key", "value")
	if err != nil {
		fmt.Println("Ошибка")
	}

	value, err := db.Get(ctx, "key")
	if err != nil {
		fmt.Println("Ошибка")
	} else {
		fmt.Println("Значение", value)
	}

	isCheck, err := db.Check(ctx, "key1")
	if err != nil {
		fmt.Println("Ошибка")
	} else {
		fmt.Println("Ключ есть", isCheck)
	}
	err = db.Delete(ctx, "key1")
	if err != nil {
		fmt.Println("Ошибка")
	}
}
