package store_redis

import (
	"InterfaceDroch/internal/model"
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"strings"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedisClient(addr string) (*Redis, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // TOCHECK не получается подключиться с паролем, только без пароля
	})
	return &Redis{client: rdb}, nil
}

func (r Redis) Set(ctx context.Context, user *model.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}
	key := "user:" + strconv.Itoa(int(user.Id)) //Создание ключа из ID и преобразование ID в строку, т.к. редиска кушает только строчки
	return r.client.WithContext(ctx).Set(key, userJSON, 20*time.Second).Err()
}

func (r Redis) Get(ctx context.Context, id int64) (*model.User, error) {
	key := "user:" + strconv.Itoa(int(id))
	userJSON, err := r.client.WithContext(ctx).Get(key).Result()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(userJSON), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r Redis) Check(ctx context.Context, id int64) (bool, error) {
	key := "user:" + strconv.Itoa(int(id))

	exists, err := r.client.WithContext(ctx).Exists(key).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
func (r Redis) Delete(ctx context.Context, id int64) error {
	key := "user:" + strconv.Itoa(int(id))

	_, err := r.client.WithContext(ctx).Del(key).Result()
	if err != nil {
		return err
	}
	return nil
}
func (r Redis) GetAllId(ctx context.Context) []int64 {
	keys, err := r.client.WithContext(ctx).Keys("user:*").Result()
	if err != nil {
		log.Printf("Не удалось получить ключи пользователей: %v", err)
		return nil
	}
	var ids []int64
	// Вот тут надо будет подучить, логика понятна, но муть мутная с этими строками
	for _, key := range keys {
		// Разделяет строки сеператором в виде двоеточия
		parts := strings.Split(key, ":")
		// Проверка есть ли 2 части. Это необходимо для того, чтобы убедиться, что ключ имеет правильный формат "user:<id>"
		if len(parts) >= 2 {
			// Если ключ разделен на две или более частей, мы извлекаем вторую часть и кладем в переменную
			idStr := parts[1]
			// Попытка преобразовать строку в число с десятичной системой и битовой длинной = 64
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err == nil {
				// Добавляем id в срез ids. Таким образом, мы накапливаем все идентификаторы пользователей из ключей, которые соответствуют шаблону "user:*"
				ids = append(ids, id)
			}
		}
	}
	return ids
}
