package main

import (
	"InterfaceDroch/internal/model"
	"InterfaceDroch/internal/store_inter"
	store_map "InterfaceDroch/internal/store_inter/instance/go"
	store_post "InterfaceDroch/internal/store_inter/instance/post"
	store_redis "InterfaceDroch/internal/store_inter/instance/redis"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Работа с файлом env
	envFilePath := "./.env"
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Не удалось загрузить переменную окружения")
	}

	StoreType := store_inter.StoreRedis

	var storage store_inter.Storage

	switch StoreType {
	case store_inter.StoreMap:
		storage = store_map.New()

	case store_inter.StorePostgres:

		// Вытаскиваем переменную
		password := os.Getenv("POSTGRES_PASSWORD")
		connectStr := fmt.Sprintf("user=postgres password=%s dbname=postgres host=localhost port=5432", password)
		//Создаем соединение
		storage, err = store_post.NewPostgresDB(connectStr)
		if err != nil {
			log.Fatalf("Ошибка создания БД: %v", err)
		}

		//TODO закрыть соединение через defer close. Я не понимаю как мне вытащить метод close из структуры в другом пакете

	case store_inter.StoreRedis:
		addr := os.Getenv("REDIS_ADDRESS")
		//password := os.Getenv("REDIS_PASSWORD") не получается подключиться к бд с паролем, только без пароля
		storage, err = store_redis.NewRedisClient(addr)
		if err != nil {
			log.Fatalf("Не удалось подключиться к редиске: %v", err)
		}

		//TODO закрыть соединение через defer close. Я не понимаю как мне вытащить метод close из структуры в другом пакете
	}
	ctx := context.Background()

	println()

	ids := storage.GetAllId(ctx)
	fmt.Printf("Список пользователей: %d\n", ids)

	// add 1 user
	u1 := model.User{
		Id:   1,
		Name: "Pon4ik",
	}
	err = storage.Set(ctx, &u1)
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

	id = 1
	user, err := storage.Get(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user)

	id = 2
	user, err = storage.Get(ctx, id)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(user)

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
