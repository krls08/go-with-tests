package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

//type DefaultSleeper struct {
//	Calls int
//}
//
//func (s *DefaultSleeper) Sleep() {
//	time.Sleep(1 * time.Second)
//}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)

}

func Countdown(out io.Writer, sleeper Sleeper) {

	//fmt.Fprintf(out, "%d", 3)

	for i := countdownStart; i > 0; i-- {
		//fmt.Fprintf(out, "%d\n", i)
		fmt.Fprintln(out, i)
		sleeper.Sleep()

	}
	fmt.Fprint(out, finalWord)
}

func main() {
	//	sleeper := &DefaultSleeper{}
	//	Countdown(os.Stdout, sleeper)
	sleeper := &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	}

	Countdown(os.Stdout, sleeper)
}
