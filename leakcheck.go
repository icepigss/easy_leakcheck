package easy_leackcheck

import (
	"context"
	"runtime"
	"time"
)

func Check(dur time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())

	timer := time.AfterFunc(dur, cancel)

	loopCheck(ctx)

	timer.Stop()

	cancel()
}

func loopCheck(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			println("has leak")
		case <-ticker.C:
			if !leak() {
				println("no leak")
				return
			}
			continue
		}
		break
	}
}

func leak() bool {
	if runtime.NumGoroutine() > 2 {
		return true
	}
	return false
}
