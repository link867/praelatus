package bolt

import (
	"encoding/json"
	"log"

	"github.com/boltdb/bolt"
	"github.com/praelatus/backend/models"
	"github.com/praelatus/backend/store"
)

// SessionStore implements store.SessionStore for a boltdb based cache
type SessionStore struct {
	db *bolt.DB
}

// Get will get the sesion information for the given session key
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

// Set will set the session information for the given session key
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

// New will open boltdb at filename for storing session info in
func New(filename string) store.SessionStore {
	ss := &SessionStore{}
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		log.Panicln("Error starting session db:", err.Error())
	}

	ss.db = db
	return ss
}
