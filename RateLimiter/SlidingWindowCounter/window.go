package main

import (
	"fmt"
	"time"
)

type SlidingWindowCounter struct {
	threshold   int
	timeframe   int
	prevCounter int
	currWindow  *Window
}

type Window struct {
	startTime time.Time
	counter   int
}

func NewSlidingWindowCounter(threshold int, timeframe int) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		threshold:   threshold,
		timeframe:   timeframe,
		prevCounter: 0,
		currWindow:  NewWindow(),
	}
}

func NewWindow() *Window {
	return &Window{
		startTime: time.Now(),
		counter:   0,
	}
}

func (w *Window) Increment() {
	w.counter++
}

func (w *Window) Counter() int {
	return w.counter
}

func (w *Window) StartTime() time.Time {
	return w.startTime
}

func (w *Window) SetStartTime(startTime time.Time) *Window {
	w.startTime = startTime
	return w
}

func (w *Window) SetCounter() *Window {
	w.counter = 0
	return w
}

func (s *SlidingWindowCounter) IsAllowed(packet *Packet) bool {
	currTime := time.Now()

	fmt.Println(time.Since(s.Window().StartTime()).Abs().Seconds(),  time.Duration(s.WindowTimeFrame()).Seconds() )
	if time.Since(s.Window().StartTime()).Seconds() > time.Duration(s.WindowTimeFrame()).Seconds() {
		s.SetPrevCounter(s.Window().Counter())
		s.Window().SetStartTime(currTime)
		s.Window().SetCounter()
	}

	// curr count + prev count * (*age of time elapsed from prev window)
	count := s.Window().Counter() + (s.PrevCounter() * (s.WindowTimeFrame() - int(currTime.Sub(s.Window().StartTime())/time.Duration(s.WindowTimeFrame()))))
	if count > s.Threshold() {
		fmt.Println("dropping packet ", packet.PacketId())
		return false
	}
	s.Window().Increment()
	return true
}

func (s *SlidingWindowCounter) SetPrevCounter(counter int) *SlidingWindowCounter {
	s.prevCounter = counter
	return s
}

func (s *SlidingWindowCounter) PrevCounter() int {
	return s.prevCounter
}

func (s *SlidingWindowCounter) Window() *Window {
	return s.currWindow
}

func (s *SlidingWindowCounter) WindowTimeFrame() int {
	return s.timeframe
}

func (s *SlidingWindowCounter) Threshold() int {
	return s.threshold
}

func (s *SlidingWindowCounter) SetThreshold(threshold int) *SlidingWindowCounter {
	s.threshold = threshold
	return s
}

func (s *SlidingWindowCounter) SetWindowTimeFrame(timeframe int) *SlidingWindowCounter {
	s.timeframe = timeframe
	return s
}

func (s *SlidingWindowCounter) SetWindow(window *Window) *SlidingWindowCounter {
	s.currWindow = window
	return s
}
