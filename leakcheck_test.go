package easy_leackcheck

import (
	"testing"
	"time"
)

func TestCheckLeak(t *testing.T) {
	defer Check(5 * time.Second)

	go func() {
		//return
		for {
			time.Sleep(time.Second)
		}
	}()
}
