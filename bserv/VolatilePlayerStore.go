package poker

type VolatilePlayerStore struct {
	scores map[string]int
}

func (v *VolatilePlayerStore) GetPlayerScore(name string) int {
	return v.scores[name]
}

func (v *VolatilePlayerStore) RecordWin(name string) {
	v.scores[name]++
}

func (v *VolatilePlayerStore) GetLeague() []Player {
	var league []Player
	for name, wins := range v.scores {
		league = append(league, Player{name, wins})
	}
	return league
}

func NewVolatilePlayerStore() *VolatilePlayerStore {
	return &VolatilePlayerStore{map[string]int{}}
}
