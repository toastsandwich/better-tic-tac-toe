package repository

import "go.etcd.io/bbolt"

var USERS = []byte("users")

type Repository struct {
	DB *bbolt.DB
}

func NewRepository(db *bbolt.DB) (*Repository, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(USERS)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &Repository{DB: db}, nil
}

func (r *Repository) InsertUser(email, user []byte) error {
	return r.DB.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket(USERS)
		return bkt.Put(email, user)
	})
}

func (r *Repository) DeleteUser(email []byte) error {
	return r.DB.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket(USERS)
		return bkt.Delete(email)
	})
}

func (r *Repository) GetUser(email []byte) ([]byte, error) {
	user := []byte{}
	err := r.DB.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket(USERS)
		user = bkt.Get(email)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
