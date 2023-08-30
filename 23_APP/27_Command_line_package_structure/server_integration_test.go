//go:build integration
// +build integration

package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrivingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, `[]`)
	defer cleanDatabase()
	anStore, err := NewFileSystemPlayerStore(database)
	assertNoError(t, err)
	server := NewPlayerServer(anStore)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")

	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []Player{
			{"Pepper", 3},
		}
		assertLeague(t, got, want)
	})

}

// INMEMORY DATABASE test
//func TestRecordingWinsAndRetrivingThem(t *testing.T) {
//	anStore := NewInMemoryPlayerStore()
//	server := NewPlayerServer(anStore)
//	player := "Pepper"
//
//	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
//	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
//	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
//
//	t.Run("get score", func(t *testing.T) {
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, newGetScoreRequest(player))
//
//		assertStatus(t, response.Code, http.StatusOK)
//
//		assertResponseBody(t, response.Body.String(), "3")
//
//	})
//
//	t.Run("get league", func(t *testing.T) {
//		response := httptest.NewRecorder()
//		server.ServeHTTP(response, newGetLeagueRequest())
//		assertStatus(t, response.Code, http.StatusOK)
//
//		got := getLeagueFromResponse(t, response.Body)
//		want := []Player{
//			{"Pepper", 3},
//		}
//		assertLeague(t, got, want)
//	})
//
//}
