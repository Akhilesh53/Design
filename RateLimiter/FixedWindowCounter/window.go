package main

import (
	"fmt"
	"time"
)

type Window struct {
	startTime time.Time
	counter   int
}

type FixedWindowCounter struct {
	threshold       int
	windowTimeFrame time.Duration
	currWindow      *Window
}

func NewWindow() *Window {
	return &Window{
		startTime: time.Now(),
		counter:   0,
	}
}

func NewFixedWindowCounter(threshold int, size time.Duration) *FixedWindowCounter {
	return &FixedWindowCounter{
		threshold:       threshold,
		windowTimeFrame: size,
		currWindow:      NewWindow(),
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

func (f *FixedWindowCounter) IsAllowed() bool {
	if time.Duration(time.Since(f.currWindow.StartTime()).Seconds()) > time.Duration(f.windowTimeFrame.Seconds()) {
		fmt.Println("window expired :: ", f.Window().StartTime().String())
		f.currWindow.SetStartTime(f.Window().StartTime().Add(f.WindowTimeFrame() * time.Second))
		f.currWindow.counter = 0
	}
	f.currWindow.Increment()
	fmt.Println("window counter :: ", f.currWindow.Counter(), f.Window().StartTime().String())
	return f.currWindow.Counter() <= f.threshold
}

func (f *FixedWindowCounter) Threshold() int {
	return f.threshold
}

func (f *FixedWindowCounter) WindowTimeFrame() time.Duration {
	return f.windowTimeFrame
}

func (f *FixedWindowCounter) Window() *Window {
	return f.currWindow
}

func (f *FixedWindowCounter) SetThreshold(threshold int) *FixedWindowCounter {
	f.threshold = threshold
	return f
}

func (f *FixedWindowCounter) SetWindowTimeFrame(tf time.Duration) *FixedWindowCounter {
	f.windowTimeFrame = tf
	return f
}

func (f *FixedWindowCounter) SetWindow(window *Window) *FixedWindowCounter {
	f.currWindow = window
	return f
}
