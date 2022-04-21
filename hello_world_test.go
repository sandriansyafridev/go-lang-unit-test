package golangunittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func DisplayName(s string) string {
	return "Hello " + s
}

type Test struct {
	Name     string
	Request  string
	Expected string
}

func TestSubTest(t *testing.T) {

	tests := []Test{
		{Name: "Sandrian", Request: "Sandrian", Expected: "Hello Sandrian"},
		{Name: "Hafid", Request: "Hafid", Expected: "Hi Hafid"},
		{Name: "Fikri", Request: "Fikri", Expected: "Hello Fikri!"},
	}

	for _, test := range tests {

		t.Run(test.Name, func(t *testing.T) {

			result := DisplayName(test.Request)

			assert.Equal(t, test.Expected, result)

		})

	}

}
