package store_new

import (
	"InterfaceDroch/internal/store"
	store_map "InterfaceDroch/internal/store/instance/go"
	store_post "InterfaceDroch/internal/store/instance/post"
)

type StoreType string

var (
	StoreMap      StoreType = "StoreMap"
	StorePostgres StoreType = "StorePostgres"
)

func NewStore(storeType StoreType) (store.Storage, error) {
	switch storeType {
	case StoreMap:
		return store_map.New(), nil
	case StorePostgres:
		return store_post.NewPostgresDB()
	}
	return nil, nil
}
