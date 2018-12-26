package poker

// Game manages the state of current poker round
type Game interface {
	Start(playersCount int)
	Finish(winner string)
}