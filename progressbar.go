package progressbar

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ProgressBar progress bar
type ProgressBar struct {
	ch           chan int64
	done         chan struct{}
	max          int64
	flushInterMs int64
}

// NewProgressBar build
func NewProgressBar(max int64) *ProgressBar {
	return &ProgressBar{
		ch:           make(chan int64),
		max:          max,
		flushInterMs: 10,
		done:         make(chan struct{}, 1),
	}
}

// Start setup gocoroutine
func (bar *ProgressBar) Start() chan<- int64 {
	go bar.execute()
	return bar.ch
}

// Close close
func (bar *ProgressBar) Close() {
	bar.done <- struct{}{}
}

func (bar *ProgressBar) execute() {
	progress := int64(0)
	ticker := time.NewTicker(time.Duration(bar.flushInterMs) * time.Millisecond)
	defer ticker.Stop()
	defer fmt.Println()
	for {
		select {
		case num := <-bar.ch:
			progress = progress + num
		case <-ticker.C:
			bar.flush(progress)
		case <-bar.done:
			return
		}
	}
}

func (bar *ProgressBar) flush(progress int64) {

	percent := float64(progress) / float64(bar.max)
	//if progress > bar.max {
	//	percent = 1
	//}

	progressBarWidth := 50
	hashesNum := int(percent * float64(progressBarWidth))
	msg := strings.Repeat("=", progressBarWidth)
	if progressBarWidth > hashesNum {
		msg = strings.Repeat("=", hashesNum)
		msg = msg + strings.Repeat(" ", progressBarWidth-hashesNum)
	}

	fmt.Printf("\x1b[0G\x1b[2K%s %d%%", msg, int(percent*100))
	os.Stdout.Sync()
}
