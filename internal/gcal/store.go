package gcal

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"go.etcd.io/bbolt"

	"google.golang.org/api/calendar/v3"
)

const (
	storeDBName = "events_gcal.db"
	existToken  = "true"
)

var (
	bucket = []byte("events")
)

func store(ctx context.Context, es []*calendar.Event) error {
	db, err := bbolt.Open(storeDBName, os.ModePerm, nil)
	if err != nil {
		return fmt.Errorf("gcal: store: %w", err)
	}
	defer db.Close()

	if err := db.Update(func(tx *bbolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return fmt.Errorf("update: %w", err)
		}

		for _, e := range es {
			j, err := e.MarshalJSON()
			if err != nil {
				return fmt.Errorf("marshal: %w", err)
			}
			dg, err := DG(j)
			if err != nil {
				return fmt.Errorf("hash: %w", err)
			}
			if err := b.Put([]byte(dg), []byte(existToken)); err != nil {
				return fmt.Errorf("put: %w", err)
			}
		}
		return nil
	}); err != nil {
		return fmt.Errorf("store: %w", err)
	}

	return nil
}

func isEventExist(_ context.Context, e *calendar.Event) (bool, error) {

	db, err := bbolt.Open(storeDBName, os.ModePerm, nil)
	if err != nil {
		return false, fmt.Errorf("gcal: store: %w", err)
	}
	defer db.Close()

	var isExist bool

	if err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucket)

		j, err := e.MarshalJSON()
		if err != nil {
			return fmt.Errorf("marshal: %w", err)
		}

		dg, err := DG(j)
		if err != nil {
			return fmt.Errorf("hash: %w", err)
		}

		if len(b.Get([]byte(dg))) > 0 {
			isExist = true
			return nil
		}

		return nil
	}); err != nil {
		return false, fmt.Errorf("get: %w", err)
	}

	return isExist, nil
}

func DG(text []byte) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write(text); err != nil {
		return "", fmt.Errorf("get digest: %w", err)
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
