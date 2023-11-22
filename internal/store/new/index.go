package store_new

import (
	"InterfaceDroch/internal/store"
	store_map "InterfaceDroch/internal/store/instance/go"
)

type StoreType string

var (
	StoreMap      StoreType
	StorePostgres StoreType
)

func NewStore(storeType StoreType) store.Storage {
	switch storeType {
	case StoreMap:
		return store_map.NewMapSore()
	case StorePostgres:
		return nil
	}
	return nil
}
