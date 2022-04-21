package golangunittest

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func HelloWorld() string {
	return "Hello World!"
}

func TestHelloWorld(t *testing.T) {

	result := HelloWorld()

	// assert.Equal(t, "Hello World", result) //assert akan memanggil fail (baris code selanjutnya akan dieksekusi)
	require.Equal(t, "Hello World", result) //require akan memanggil fatal (baris code selanjutnya tidak dieksekusi)

	log.Println("RUN APPLICATION")
	log.Println("RUN APPLICATION")
	log.Println("RUN APPLICATION")
	log.Println("RUN APPLICATION")

}
