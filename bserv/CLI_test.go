package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/0xpsilocybe/learning-go/bserv"
)

var (
	dummyBlindAlerter = &poker.SpyBlindAlerter{}
	dummyPlayerStore  = &poker.StubPlayerStore{}
	dummyStdIn        = &bytes.Buffer{}
	dummyStdOut       = &bytes.Buffer{}
)

func userSends(messages... string) io.Reader {
	return strings.NewReader(strings.Join(messages, "\n"))
}

func TestCLI(t *testing.T) {

	t.Run("start game with 3 players and finish it with 'Chris' as a winner", func(t *testing.T) {
		in := userSends("3", "Chris wins")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("start game with 8 players and finish it with 'Cleo' as a winner", func(t *testing.T) {
		in := userSends("8", "Cleo wins")
		playerStore := &poker.StubPlayerStore{}
		game := poker.NewGame(dummyBlindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing blind values", func(t *testing.T) {
		// Arrange
		in := strings.NewReader("5\nBob wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, playerStore)
		cli := poker.NewCLI(in, dummyStdOut, game)
		// Act
		cli.PlayPoker()
		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				// Assert
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

	t.Run("it prompts the user for the number of players", func(t *testing.T) {
		// Arrange
		stdOut := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewGame(blindAlerter, dummyPlayerStore)
		cli := poker.NewCLI(in, stdOut, game)
		// Act
		cli.PlayPoker()
		// Assert
		got := stdOut.String()
		want := poker.PlayerPrompt
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
		cases := []poker.ScheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		for i, want := range cases {
			t.Run(fmt.Sprint(want), func(t *testing.T) {
				if len(blindAlerter.Alerts) <= i {
					t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
				}
				got := blindAlerter.Alerts[i]
				assertScheduledAlert(t, got, want)
			})
		}
	})

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
