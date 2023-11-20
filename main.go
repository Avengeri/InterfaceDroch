package main

import (
	database "InterfaceDroch/interface/go"
	"context"
	"fmt"
)

func main() {
	db := database.Database{}
	ctx := context.Background()

	err := db.Set(ctx, "Имя", "Андрей")
	if err != nil {
		fmt.Println("Ошибка")
	}

	value := db.Data["Имя"]
	fmt.Println("Имя", value) //Нифига не выведет потому что интерфейс, а что туда класть я хз
}
