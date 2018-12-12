package poker_test

import (
	"strings"
	"testing"

	"github.com/bartek/learning-go/bserv"
)

var dummySpyAlerter = &poker.SpyBlindAlerter{}

func TestCLI(t *testing.T) {

	t.Run("record Chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record Cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in, dummySpyAlerter)
		cli.PlayPoker()
		poker.AssertPlayerWin(t, playerStore, "Cleo")
	})

	t.Run("it schedules printing blind values", func(t *testing.T) {
		in := strings.NewReader("Bob wins\n")
		playerStore := &poker.StubPlayerStore{}
		blindAlerter := &poker.SpyBlindAlerter{}
		cli := poker.NewCLI(playerStore, in, blindAlerter)
		cli.PlayPoker()
		if len(blindAlerter.Alerts) != 1 {
			t.Fatal("expected a blind alert to be scheduled")
		}
	})

}
