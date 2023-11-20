package database

import (
	"context"
	"fmt"
	"log"
)

type Database struct {
	Data map[string]interface{}
}

func (d Database) Set(ctx context.Context, key string, value interface{}) error {
	if d.Data == nil {
		d.Data = make(map[string]interface{})
	}
	d.Data[key] = value
	return nil
}

func (d Database) Get(ctx context.Context, key string) (interface{}, error) {
	if d.Data == nil {
		log.Println("Не удалось получить данные")
	}
	value, found := d.Data[key]
	if !found {
		return nil, fmt.Errorf("Ключ %s не найден в мапе", key)
	}
	return value, nil
}

func (d Database) Delete(ctx context.Context, key string) error {
	if d.Data == nil {
		log.Println("Не удалось получить данные")
	}
	_, found := d.Data[key]
	if !found {
		return fmt.Errorf("Ключ %s не найден в мапе", key)
	}
	delete(d.Data, key)
	return nil
}

func (d Database) Check(ctx context.Context, key string) (bool, error) {
	if d.Data == nil {
		log.Println("Не удалось получить данные")
	}
	_, found := d.Data[key]
	return found, nil
}

func NewStorage(key, value string) interface{} {

}
