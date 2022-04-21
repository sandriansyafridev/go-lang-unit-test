package golangunittest

import "testing"

func Hello(name string) string {
	return "Hello " + name

}

func BenchmarkHello(b *testing.B) {

	b.Run("Sandrian", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Hello("rian")
		}
	})

	b.Run("Hafid", func(b *testing.B) {

		for i := 0; i < b.N; i++ {
			Hello("Hafid")
		}
	})

}
