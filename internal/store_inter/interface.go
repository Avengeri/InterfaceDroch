package store_inter

import (
	"InterfaceDroch/internal/model"
	"context"
)

type StoreType string

var (
	StoreMap      StoreType = "StoreMap"
	StorePostgres StoreType = "StorePostgres"
	StoreRedis    StoreType = "StoreRedis"
)

type Storage interface {
	Set(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Check(ctx context.Context, id int64) (bool, error)
	Delete(ctx context.Context, id int64) error
	GetAllId(ctx context.Context) []int64
}
