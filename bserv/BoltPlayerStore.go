package poker

import (
	"encoding/binary"
	"fmt"
	"sort"
	"time"

	"github.com/boltdb/bolt"
)

var (
	ScoresBucket = []byte("Scores")
)

type BoltPlayerStore struct {
	db *bolt.DB
}

func (b *BoltPlayerStore) GetPlayerScore(name string) int {
	score := 0
	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ScoresBucket)
		v := b.Get([]byte(name))
		if v != nil {
			score = int(binary.BigEndian.Uint32(v))
		}
		return nil
	})
	if err != nil {
		fmt.Errorf("get player score: %s", err)
	}
	return score
}

func (b *BoltPlayerStore) RecordWin(name string) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		score := uint32(b.GetPlayerScore(name) + 1)
		scoreBuffer := make([]byte, 4)
		binary.BigEndian.PutUint32(scoreBuffer, score)
		b := tx.Bucket(ScoresBucket)
		err := b.Put([]byte(name), scoreBuffer)
		return err
	})
	if err != nil {
		fmt.Errorf("record win: %s", err)
	}
}

func (b *BoltPlayerStore) GetLeague() []Player {
	var league []Player
	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(ScoresBucket)
		b.ForEach(func(k, v []byte) error {
			name := string(k)
			wins := int(binary.BigEndian.Uint32(v))
			league = append(league, Player{name, wins})
			return nil
		})
		return nil
	})
	if err != nil {
		return nil
	}
	sort.Slice(league, func(i, j int) bool {
		return league[i].Wins > league[j].Wins
	})
	return league
}

func (b *BoltPlayerStore) Close() {
	b.db.Close()
}

func NewBoltPlayerStore(path string) (store *BoltPlayerStore, err error) {
	db, err := bolt.Open(
		path,
		0600,
		&bolt.Options{Timeout: 1 * time.Second},
	)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(ScoresBucket)
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &BoltPlayerStore{db}, nil
}
