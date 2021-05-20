package main

import (
	"fmt"
	"testing"
  "strconv"
)

func main() {
	result1 := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%.6fs", 0.1111)
		}
	})
	fmt.Printf("%s\n", result1)
	result2 := testing.Benchmark(func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = strconv.FormatFloat(0.1111, 'f', 9, 64) + "s"
		}
	})
	fmt.Printf("%s\n", result2)

}

