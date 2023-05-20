package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("return fast url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer10(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("returns an error if a server does not respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Millisecond)

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected error but didn't get one")
		}

	})

}

func BenchmarkRacer(b *testing.B) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0)

	defer slowServer.Close()
	defer fastServer.Close()

	b.Run("fast url is returned", func(b *testing.B) {
		for i := 0; i < 1000; i++ {
			_, _ = Racer10(slowServer.URL, fastServer.URL)
		}

	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
