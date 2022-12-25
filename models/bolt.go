package models

import (
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

func Open(dbName string) (*bolt.DB, error) {
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	return db, nil
}
