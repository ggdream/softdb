package pool

import (
	"sync"
	"testing"
	"time"
)


func TestNew(t *testing.T) {
	var wg sync.WaitGroup
	p := New(3)
	for i := 0; i < 18; i++ {
		wg.Add(1)
		a := func(s int) func() {
			return func() {
				defer wg.Done()
				println(s)
			time.Sleep(time.Second * 1)
			}
		}(i)
		p.Add(a)
	}
	p.Run()
	wg.Wait()
}
