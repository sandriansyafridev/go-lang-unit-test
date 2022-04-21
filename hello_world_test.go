package golangunittest

import (
	"log"
	"testing"
)

func HelloWorld() string {
	return "Hello World!"
}

func TestHelloWorld(t *testing.T) {
	if result := HelloWorld(); result != "Hello World" {
		// t.Error("Return value must be 'Hello World!'") // jika test gagal atau fail maka baris code selanjutnya akan tetap dieksekusi
		t.Fatal("Return value must be 'Hello World'") //jika test gagal atau fail maka baris code selanjutya tidak akan dieksekusi
	}

	log.Println("RUNNING APP")
	log.Println("RUNNING APP")
	log.Println("RUNNING APP")

	log.Println("SUCCESS")
}
