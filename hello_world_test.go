package golangunittest

import (
	"log"
	"testing"
)

func HelloWorld() string {
	return "Hello World!"
}

func TestHelloWorld(t *testing.T) {

	result := HelloWorld()
	if result != "Hello World!" {
		panic("ERROR")
	}

	log.Println("SUCCESS")

}
