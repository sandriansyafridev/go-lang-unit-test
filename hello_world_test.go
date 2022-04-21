package golangunittest

import "testing"

func Hello(name string) string {
	return "Hello " + name

}

type Tests struct {
	Name    string
	Request string
}

func BenchmarkHello(b *testing.B) {

	tests := []Tests{
		{
			Name:    "Sandrian",
			Request: "sandrian",
		},
		{
			Name:    "Hafid",
			Request: "hafid",
		},
	}

	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(test.Request)
			}
		})
	}

}
