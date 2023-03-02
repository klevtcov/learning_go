package ma0001

import (
	"fmt"
	"strings"
	"time"
)

type Stopwatch struct {
	startTime time.Time
	results   []time.Duration
}

func (s *Stopwatch) Start() {
	s.startTime = time.Now()
}

func (s *Stopwatch) SaveSplit() {
	delta := time.Since(s.startTime)
	s.results = append(s.results, delta)
}

func (s *Stopwatch) GetResults() string {
	var stringDuration []string
	for _, val := range s.results {
		stringDuration = append(stringDuration, fmt.Sprint(val.Seconds()))
	}
	result := strings.Join(stringDuration, "s ")
	result += "s"
	return result
}

func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

	// [1.004643466s 1.505609247s 1.806269902s]
}

// Реализуйте тип Stopwatch, который будет сохранять время каждой промежуточной фиксации секундомера и выдавать результаты относительно общего времени старта.
// Тип должен обладать следующими методами:
// Start() — запустить/сбросить секундомер;
// SaveSplit() — сохранить промежуточное время;
// GetResults() []time.Duration — вернуть текущие результаты.
//
//
//

// package main

// import (
//     "fmt"
//     "time"
// )

// type Stopwatch struct {
//     startTime time.Time
//     splits    []time.Time
// }

// func (sw *Stopwatch) Start() {
//     sw.startTime = time.Now()
//     sw.splits = nil
// }

// func (sw *Stopwatch) SaveSplit() {
//     sw.splits = append(sw.splits, time.Now())
// }

// func (sw Stopwatch) GetResults() (retResults []time.Duration) {
//     for _, splitTime := range sw.splits {
//         retResults = append(retResults, splitTime.Sub(sw.startTime))
//     }

//     return
// }

// func main() {
//     sw := Stopwatch{}
//     sw.Start()

//     time.Sleep(1 * time.Second)
//     sw.SaveSplit()

//     time.Sleep(500 * time.Millisecond)
//     sw.SaveSplit()

//     time.Sleep(300 * time.Millisecond)
//     sw.SaveSplit()

//     fmt.Println(sw.GetResults())
// }
