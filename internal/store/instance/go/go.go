package store_map

import (
	"InterfaceDroch/internal/model"
	"InterfaceDroch/internal/store"
	"context"
	"errors"
	"fmt"
)

var _ store.Storage = &database{}

type database struct {
	data map[int64]*model.User
}

func NewMapSore() store.Storage {
	return &database{
		data: make(map[int64]*model.User),
	}
}

func (d database) Set(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.New("user is empty")
	}
	d.data[user.Id] = user
	return nil
}

func (d database) Get(ctx context.Context, id int64) (*model.User, error) {
	value, ok := d.data[id]
	if !ok {
		return nil, fmt.Errorf("Ключ %d не найден в мапе", id)
	}
	return value, nil
}

func (d database) Check(ctx context.Context, id int64) (bool, error) {
	_, ok := d.data[id]
	return ok, nil
}

func (d database) Delete(ctx context.Context, id int64) error {
	delete(d.data, id)
	return nil
}

func (d database) GetAllId(ctx context.Context) []int64 {
	var ids []int64

	for key, _ := range d.data {
		ids = append(ids, key)
	}
	return ids
}
