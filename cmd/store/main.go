package main

import (
	"InterfaceDroch/internal/model"
	store_new "InterfaceDroch/internal/store/new"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()
	storage, _ := store_new.NewStore(store_new.StorePostgres)
	if storage == nil {
		log.Fatal("storage is nil")
		return
	}
	println()
	ids := storage.GetAllId(ctx)
	fmt.Printf("Список пользователей: %d\n", ids)

	// add 1 user
	u1 := model.User{
		Id:   1,
		Name: "Pon4ik",
	}
	err := storage.Set(ctx, &u1)
	if err != nil {
		log.Fatal(err)
		return
	}
	// add 2 user
	u2 := model.User{
		Id:   2,
		Name: "Gofer",
	}
	err = storage.Set(ctx, &u2)
	if err != nil {
		log.Fatal(err)
		return
	}
	// get all users
	ids = storage.GetAllId(ctx)
	fmt.Printf("Список пользователей: %d\n", ids)

	// check 1 user
	id := int64(1)
	ok, err := storage.Check(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	str := "ЕСТЬ"
	if !ok {
		str = "НЕТУ"
	}
	fmt.Printf("Пользователь %d %s в системе\n", id, str)

	// delete 1 user
	err = storage.Delete(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	// get all users
	ids = storage.GetAllId(ctx)
	fmt.Printf("Список пользователей: %d\n", ids)
}
