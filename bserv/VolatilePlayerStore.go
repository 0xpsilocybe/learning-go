package main

type VolatilePlayerStore struct {
	scores map[string]int
}

func (v *VolatilePlayerStore) GetPlayerScore(name string) int {
	return v.scores[name]
}

func (v *VolatilePlayerStore) RecordWin(name string) {
	v.scores[name]++
}

func NewVolatilePlayerStore() *VolatilePlayerStore {
	return &VolatilePlayerStore{map[string]int{}}
}
