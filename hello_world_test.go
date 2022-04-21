package golangunittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func DisplayName(s string) string {
	return "Hello " + s
}

func TestSubTest(t *testing.T) {

	t.Run("Sandrian", func(t *testing.T) {
		result := DisplayName("Sandrian")

		assert.Equal(t, "Hello Sandrian!", result)
	})

	t.Run("Hafid", func(t *testing.T) {
		result := DisplayName("Hafid")

		assert.Equal(t, "Hello Hafid", result)
	})

}
