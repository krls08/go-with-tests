package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloCustom(t *testing.T) {
	t.Run("say hello to krls", func(t *testing.T) {
		got := HelloC("krls")
		want := "Hello, krls!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello, world, when and empty string is supplied", func(t *testing.T) {
		got := HelloC("")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
		//		if got != want {
		//			t.Errorf("got %q, want %q", got, want)
		//		}
	})
}

func TestHelloLang(t *testing.T) {
	t.Run("in spanisht", func(t *testing.T) {
		got := HelloLang("Carlos", "Spanish")
		want := "Hola, Carlos!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in french", func(t *testing.T) {
		got := HelloLang("Jeremy", "French")
		want := "Bonjour, Jeremy!"
		assertCorrectMessage(t, got, want)
	})
}
func TestGetPrefix(t *testing.T) {
	t.Run("get spanish prefix", func(t *testing.T) {
		got := getPrefix("Spanish")
		want := "Hola"
		assertCorrectMessage(t, got, want)
	})
	t.Run("get french prefix", func(t *testing.T) {
		got := getPrefix("French")
		want := "Bonjour"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
