package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(aDuration time.Duration) {
	s.durationSlept = aDuration
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("configure duration", func(t *testing.T) {
		timeToSleep := 5 * time.Second

		spyTime := &SpyTime{}
		sleeper := &ConfigurableSleeper{
			duration: timeToSleep,
			sleep:    spyTime.Sleep,
		}

		sleeper.Sleep()

		if spyTime.durationSlept != timeToSleep {
			t.Errorf("should have slept for %v, but slept %v", timeToSleep, spyTime.durationSlept)
		}

	})
}

func TestMocking(t *testing.T) {

	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := new(SpyCountdownOperations)

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("Sleep before each write", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, but got %v", want, spySleepPrinter.Calls)
		}

	})

}
