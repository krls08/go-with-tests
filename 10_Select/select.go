package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer10(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, 10*time.Second)

}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("Timeout waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {

	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}

func main() {
	fmt.Println("start main")
}

//func Racer(url0, url1 string) (anUrl string, err error) {
//	duration0, err0 := measureResponseTime(url0)
//	duration1, err1 := measureResponseTime(url1)
//	_, _ = err0, err1
//
//	if duration0 > duration1 {
//		return url1, nil
//	}
//	return url0, nil
//}
//
//func measureResponseTime(url string) (time.Duration, error) {
//	startTime := time.Now()
//	_, err := http.Get(url)
//	if err != nil {
//		return 0, err
//	}
//	duration := time.Since(startTime)
//	return duration, nil
//}
