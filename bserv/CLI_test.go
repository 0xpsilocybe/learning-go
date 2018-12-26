package poker_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/0xpsilocybe/learning-go/bserv"
)

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

type GameSpy struct {
	StartCalled bool
	StartCalledWith int
	FinishCalled bool
	FinishCalledWith string
}

func (g *GameSpy) Start(playerCount int) {
	g.StartCalled = true
	g.StartCalledWith = playerCount
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalled = true
	g.FinishCalledWith = winner
}

func userSends(messages... string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		// Arrange
		game := &GameSpy{}
		in := userSends("whatever")
		out := &bytes.Buffer{}
		cli := poker.NewCLI(in, out, game)
		// Act
		cli.PlayPoker()
		// Assert
		wantPrompt := poker.PlayerPrompt + poker.BadPlayerInputErrorMessage
		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, out, wantPrompt)
	})

	t.Run("it prints an error when a winner is announced in a wrong way and does not finish the game", func(t *testing.T) {
		// Arrange
		game := &GameSpy{}
		in := userSends("6", "Lloyd is a champ")
		out := &bytes.Buffer{}
		cli := poker.NewCLI(in, out, game)
		// Act
		cli.PlayPoker()
		// Assert
		wantPrompt := poker.PlayerPrompt + poker.BadWinnerAnnouncementErrorMessage
		assertMessagesSentToUser(t, out, wantPrompt)
		assertGameStartedWith(t, game, 6)
		assertGameNotFinished(t, game)
	})

	t.Run("start game with 3 players and finish it with 'Chris' as a winner", func(t *testing.T) {
		// Arrange
		game := &GameSpy{}
		in := userSends("3", "Chris wins")
		out := &bytes.Buffer{}
		cli := poker.NewCLI(in, out, game)
		// Act
		cli.PlayPoker()
		// Assert
		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertGameFinishedWith(t, game, "Chris")
	})

	t.Run("start game with 8 players and finish it with 'Cleo' as a winner", func(t *testing.T) {
		// Arrange
		game := &GameSpy{}
		in := userSends("8", "Cleo wins")
		out := &bytes.Buffer{}
		cli := poker.NewCLI(in, out, game)
		// Act
		cli.PlayPoker()
		// Assert
		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 8)
		assertGameFinishedWith(t, game, "Cleo")
	})

}

func assertGameStartedWith(t *testing.T, game *GameSpy, playerCount int) {
	t.Helper()
	if game.StartCalledWith != playerCount {
		t.Errorf(
			"wanted Start called with %d, but got %d",
			playerCount,
			game.StartCalledWith,
		)
	}
}

func assertGameFinishedWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishCalledWith != winner {
		t.Errorf(
			"wanted Finish called with %s, but got %s",
			winner,
			game.FinishCalledWith,
		)
	}
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Error("game should not have started")
	}
}

func assertGameNotFinished(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.FinishCalled {
		t.Error("game should not have finished")
	}
}

func assertScheduledAlert(t *testing.T, got, want poker.ScheduledAlert) {
	t.Helper()
	if got.Amount != want.Amount {
		t.Errorf("got amount %d, want %d", got.Amount, want.Amount)
	}
	if got.At != want.At {
		t.Errorf("got scheduled time of %v, want %v", got.At, want.At)
	}
}

func assertMessagesSentToUser(t *testing.T, out *bytes.Buffer, messages... string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := out.String()
	if got != want {
		t.Errorf("'%s' was sent to user, want '%+v'", got, want)
	}
}