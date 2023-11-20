package interfaces

import "context"

type Inter interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
	Check(ctx context.Context, key string) (bool, error)
}
