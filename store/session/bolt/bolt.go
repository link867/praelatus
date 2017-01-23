package bolt

import (
	"encoding/json"

	"github.com/boltdb/bolt"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

// SessionStore implements store.SessionStore for a boltdb based cache
type SessionStore struct {
	db *bolt.DB
}

func (c *SessionStore) Get(key string) (models.User, error) {
	var u models.User
	var jsn []byte

	c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("sessions"))
		jsn = b.Get([]byte(key))
		return nil
	})

	if jsn == nil {
		return u, store.ErrNoSession
	}

	err := json.Unmarshal(jsn, &u)
	return u, err
}

func (c *SessionStore) Set(key string, model models.User) error {
	jsn, err := json.Marshal(model)
	if err != nil {
		return err
	}

	return c.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("sessions"))
		if err != nil {
			return err
		}

		return bucket.Put([]byte(key), jsn)
	})
}

func New(filename string) (*SessionStore, error) {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}

	ss := SessionStore{db}
	return &ss, err
}
