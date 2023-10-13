package progressbar_test

import (
	"testing"
	"time"
)

func TestBar(t *testing.T) {

	bar := pkg.NewProgressBar(100)

	ch := bar.Start()

	for i := int64(1); i < 100; i++ {
		ch <- 1
		time.Sleep(100 * time.Millisecond)
	}

}
