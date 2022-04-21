package golangunittest

import "testing"

func Hello(name string) string {
	return "Hello " + name

}

func BenchmarkHelloRian(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Hello("rian")
	}

}

func BenchmarkHelloSandrian(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Hello("SANDRIAN SYAFRI")
	}

}
