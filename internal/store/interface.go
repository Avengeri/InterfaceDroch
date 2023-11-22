package store

import (
	"InterfaceDroch/internal/model"
	"context"
)

type Storage interface {
	Set(ctx context.Context, user *model.User) error
	Get(ctx context.Context, id int64) (*model.User, error)
	Check(ctx context.Context, id int64) (bool, error)
	Delete(ctx context.Context, id int64) error
	GetAllId(ctx context.Context) []int64
}
