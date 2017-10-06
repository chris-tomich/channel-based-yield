package channelbasedyield

import (
	"testing"
)

func iterateItemsWithYield() chan int {
	c := make(chan int)

	go (func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	})()

	return c
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
		sum := 0

		for i := 0; i < 1000; i++ {
			sum += i
		}
	}
}
