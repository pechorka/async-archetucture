package ebadger

import (
	"testing"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/stretchr/testify/require"
)

func Test_setGet(t *testing.T) {
	db, teardown := prepareDB(t)
	t.Cleanup(teardown)
	so := require.New(t)

	type data struct {
		ID    string    `json:"id"`
		Date  time.Time `json:""`
		Flag  bool      `json:"flag"`
		Float float64   `json:"float"`
		Int   int       `json:"int"`
	}

	someData := data{
		ID:    "bla-bla",
		Date:  time.Date(2021, time.November, 8, 8, 4, 30, 0, time.UTC),
		Flag:  true,
		Float: 36.6,
		Int:   42,
	}

	err := db.Update(func(txn *badger.Txn) error {
		return SetMarshal(txn, []byte(someData.ID), someData)
	})
	so.NoError(err)

	var retrivedData data
	err = db.View(func(txn *badger.Txn) error {
		return GetUnmarshal(txn, []byte(someData.ID), &retrivedData)
	})
	so.NoError(err)

	so.Equal(someData, retrivedData)
}

func prepareDB(t *testing.T) (*badger.DB, func()) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		t.Fatal(err)
	}

	teardown := func() {
		err := db.Close()
		if err != nil {
			t.Errorf("erro closing db %v", err)
		}
	}

	return db, teardown
}
