package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
	write          = "write"
	sleep          = "sleep"
)

type Sleeper interface {
	Sleep()
}

// type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}
type SpySleeper struct {
	Calls int
}
type SpyTime struct {
	durationSlept time.Duration
}
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, "%s", finalWord)
}
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
