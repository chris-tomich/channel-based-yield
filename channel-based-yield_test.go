package channelbasedyield

import (
	"testing"
)

func iterateItemsWithYield(finish ...chan interface{}) chan int {
	c := make(chan int, 50)

	if len(finish) == 0 {
		go (func() {
			defer close(c)
			for i := 0; i < 1000000; i++ {
				c <- i
			}
		})()
	} else {
		go (func() {
			defer close(c)
			for i := 0; i < 1000000; i++ {
				select {
				case <-finish[0]:
					return
				case c <- i:
				}
			}
		})()
	}

	return c
}

func iterateItemsWithoutYield() []int {
	ii := make([]int, 0, 1000000)

	for i := 0; i < 1000000; i++ {
		ii = append(ii, i)
	}

	return ii
}

// BenchmarkChannelYield benchmarks the pattern for using a channel like a C# yield.
func BenchmarkChannelYield(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := iterateItemsWithYield()

		sum := 0

		for myNum := range c {
			sum += myNum
		}
	}
}

// BenchmarkStandardLoop is just a standard loop.
func BenchmarkStandardLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ii := iterateItemsWithoutYield()

		sum := 0

		for myNum := range ii {
			sum += myNum
		}
	}
}

// BenchmarkChannelYieldEarlyFinish benchmarks the pattern for using a channel like a C# yield. It stops at the 10000th item.
func BenchmarkChannelYieldEarlyFinish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		finish := make(chan interface{})
		c := iterateItemsWithYield(finish)

		sum := 0

		for myNum := range c {
			sum += myNum

			if myNum == 10000 {
				close(finish)
				break
			}
		}
	}
}

// BenchmarkStandardLoopEarlyFinish is just a standard loop. It stops at the 10000th item.
func BenchmarkStandardLoopEarlyFinish(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ii := iterateItemsWithoutYield()

		sum := 0

		for myNum := range ii {
			sum += myNum

			if myNum == 10000 {
				break
			}
		}
	}
}
