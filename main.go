package main

import (
	"InterfaceDroch/interface/post"
	"context"
	"fmt"
	"log"
)

func main() {

	//db := database.NewData()
	//
	//ctx := context.Background()
	//
	//err := db.Set(ctx, "key", "value")
	//if err != nil {
	//	fmt.Println("Ошибка")
	//}
	//
	//value, err := db.Get(ctx, "key")
	//if err != nil {
	//	fmt.Println("Ошибка")
	//} else {
	//	fmt.Println("Значение", value)
	//}
	//
	//isCheck, err := db.Check(ctx, "key1")
	//if err != nil {
	//	fmt.Println("Ошибка")
	//} else {
	//	fmt.Println("Ключ есть", isCheck)
	//}
	//err = db.Delete(ctx, "key1")
	//if err != nil {
	//	fmt.Println("Ошибка")
	//}

	ctx := context.Background()
	conn, err := post.NewTableStorage(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close(ctx) //Обработать ошибку

	fmt.Println("Таблица 'users' успешно создана.")

	u := post.User{}

	id := int64(1)
	name := "Andrey"
	err = u.SetStorage(ctx, id, name)
	if err != nil {
		log.Println("Не удалось создать пользователя", err)
	}
	fmt.Printf("Пользователь %s успешно добавлен\n", name)

	//тут будет работать корректно, если будет указан id который есть в бд, тк это должен быть метод, и аргументы должны вытягиваться из структуры, а не вводится вручную
	id = 1

	userExists, err := u.CheckStorage(ctx, id)
	if err != nil {
		log.Println("Ошибка проверки пользователя", err)
	}

	if userExists {
		fmt.Println("Пользователь существует.")
	} else {
		fmt.Println("Пользователя нет в бд")
	}

	id = 1

	name, err = u.GetStorage(ctx, id)
	if err != nil {
		log.Println("Не удалось получить информацию о пользователе", err)
	}
	fmt.Println(name)

	id = 2

	err = u.DeleteStorage(ctx, id)
	if err != nil {
		log.Println("Не удалось удалить пользователя", err)
	}

}
