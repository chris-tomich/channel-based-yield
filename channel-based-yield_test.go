package channelbasedyield

import (
	"testing"
)

func iterateItemsWithYield() chan int {
	c := make(chan int, 1000)

	go (func() {
		for i := 0; i < 1000; i++ {
			c <- i
		}
		close(c)
	})()

	return c
}

func iterateItemsWithoutYield() []int {
	ii := make([]int, 0, 1000)

	for i := 0; i < 1000; i++ {
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
