package repository

import (
	"fmt"

	"go.etcd.io/bbolt"
)

var (
	USERS          = []byte("users")
	BLACKLISTTOKEN = []byte("blacklisted_token")
)

type Repository struct {
	DB *bbolt.DB
}

func NewRepository(db *bbolt.DB) (*Repository, error) {
	err := db.Update(func(tx *bbolt.Tx) (err error) {
		_, err = tx.CreateBucketIfNotExists(USERS)
		if err != nil {
			return
		}
		_, err = tx.CreateBucketIfNotExists(BLACKLISTTOKEN)
		if err != nil {
			return
		}
		return
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

func (r *Repository) AddTokenToBlackList(token []byte, time []byte) error {
	return r.DB.Update(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket(BLACKLISTTOKEN)
		return bkt.Put(token, time)
	})
}

func (r *Repository) FindToken(token []byte) error {
	return r.DB.View(func(tx *bbolt.Tx) error {
		bkt := tx.Bucket(BLACKLISTTOKEN)
		if dump := bkt.Get(token); dump != nil {
			return fmt.Errorf("token is blacklisted")
		} else {
			return nil
		}
	})
}
