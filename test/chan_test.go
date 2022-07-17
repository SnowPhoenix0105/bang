package test

import (
	"testing"
	"time"
)

func TestChanClose(t *testing.T) {
	ch1 := make(chan int, 5)
	ch2 := make(chan struct{})

	go func() {
		write := func(num int) {
			ch1 <- num
			t.Logf("ch1 <- %d", num)
		}
		for i := 0; i < 5; i++ {
			write(i * 2)
			write(i*2 + 1)
			time.Sleep(1 * time.Millisecond)
		}
		close(ch1)
		t.Log("close(ch1)")
	}()

	go func() {
		for {
			num, opened := <-ch1
			if !opened {
				t.Log("ch1 closed")
				ch2 <- struct{}{}
				return
			}
			t.Logf("%d <- ch1", num)
			time.Sleep(1 * time.Millisecond)
		}
	}()

	<-ch2
}
