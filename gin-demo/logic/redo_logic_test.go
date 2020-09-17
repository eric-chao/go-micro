package service

import (
	"sync"
	"sync/atomic"
	"testing"
)

const CONCURRENT = 10

func TestRedo(t *testing.T) {
	//waitGroup := sync.WaitGroup{}
	waitGroup2 := sync.WaitGroup{}

	rpcChans := make(chan *int, CONCURRENT)
	var ops uint64
	for i := 0; i < CONCURRENT; i++ {
		waitGroup2.Add(1)
		go func() {
			for result := range rpcChans {
				//waitGroup.Done()
				atomic.AddUint64(&ops, 1)
				t.Logf("[result] %d -- %d", atomic.LoadUint64(&ops), *result)
			}
			waitGroup2.Done()
		}()
	}

	var arr = [CONCURRENT * 10]struct{}{}
	for index, _ := range arr {
		i := index
		//waitGroup.Add(1)
		rpcChans <- &i
	}

	//waitGroup.Wait()
	close(rpcChans)
	waitGroup2.Wait()

	t.Logf("ops = %d", ops)
}

func TestCount(t *testing.T) {
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				atomic.AddUint64(&ops, 1)
				t.Logf("ops = %d", ops)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	t.Logf("ops = %d", ops)
}
