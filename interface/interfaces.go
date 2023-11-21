package interfaces

import (
	"context"
)

type Inter interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Check(ctx context.Context, key string) (bool, error)
}

type InterStorage interface {
	SetStorage(ctx context.Context, id int64, name string) error
	CheckStorage(ctx context.Context, id int64) (bool, error)
	GetStorage(ctx context.Context, id int64) (string, error)
	DeleteStorage(ctx context.Context, id int64) error
}
