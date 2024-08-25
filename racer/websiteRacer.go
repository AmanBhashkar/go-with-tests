package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {

	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)
	// if aDuration < bDuration {
	// 	return a
	// }
	// return b
	return Configurableracer(a, b, tenSecondTimeout)

}
func Configurableracer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
	}()
	return ch
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }
