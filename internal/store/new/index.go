package store_new

import (
	"InterfaceDroch/internal/store"
	store_map "InterfaceDroch/internal/store/instance/go"
)

type StoreType string

var (
	StoreMap      StoreType = "StoreMap"
	StorePostgres StoreType = "StorePostgres"
)

func NewStore(storeType StoreType) store.Storage {
	switch storeType {
	case StoreMap:
		return store_map.New()
	case StorePostgres:
		return nil
	}
	return nil
}
