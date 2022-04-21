package golangunittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
)

func HelloWorld() string {
	return "Hello World!"
}

func TestSkipTest(t *testing.T) {

	if condition := true; condition {
		t.Skip()
	}

	result := HelloWorld()

	assert.Equal(t, "Hello World!", result)
}

func TestNextTest(t *testing.T) {

	result := HelloWorld()

	assert.Equal(t, "Hello World!", result)
}
