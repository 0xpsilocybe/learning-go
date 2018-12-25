package poker_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/0xpsilocybe/learning-go/bserv"
)

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := poker.NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Pepper")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		poker.AssertResponseStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Floyd")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		poker.AssertResponseStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := poker.NewGetScoreRequest("Nonamer")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		poker.AssertResponseStatus(t, response.Code, http.StatusNotFound)
	})

}

func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := poker.NewPlayerServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := poker.NewPostWinRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		poker.AssertResponseStatus(t, response.Code, http.StatusAccepted)
		if len(store.WinCalls) != 1 {
			t.Errorf("got %d calls to RecordWin, expected %d", len(store.WinCalls), 1)
		}
		if store.WinCalls[0] != player {
			t.Errorf("did not store correct winner, got '%s' want %s", store.WinCalls[0], player)
		}
	})

}

func TestLeague(t *testing.T) {

	t.Run("it returns league table as JSON on /league", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 20},
			{"Adam", 32},
			{"Voytek", 21},
		}
		store := poker.StubPlayerStore{nil, nil, wantedLeague}
		server := poker.NewPlayerServer(&store)
		request := poker.NewGetLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		poker.AssertResponseStatus(t, response.Code, http.StatusOK)
		poker.AssertContentType(t, response, poker.JsonContentType)
		got := poker.GetLeagueFromResponse(t, response.Body)
		poker.AssertLeague(t, got, wantedLeague)
	})

}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	const dbPath = "test.db"
	store, err := poker.NewBoltPlayerStore(dbPath)
	if err != nil {
		t.Fatalf("couldn't initialize store: %s", err)
	}
	server := poker.NewPlayerServer(store)
	player := "Pepper"
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetScoreRequest(player))
		poker.AssertResponseStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, poker.NewGetLeagueRequest())
		poker.AssertResponseStatus(t, response.Code, http.StatusOK)
		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})

	os.Remove(dbPath)
}
