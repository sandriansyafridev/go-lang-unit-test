package golangunittest

import (
	"log"
	"testing"
)

func HelloWorld() string {
	return "Hello World!"
}

func TestHelloWorld(t *testing.T) {
	if result := HelloWorld(); result != "Hello World!" {
		t.Fail() // jika test gagal atau fail maka baris code selanjutnya akan tetap dieksekusi
		// t.FailNow() //jika test gagal atau fail maka baris code selanjutya tidak akan dieksekusi
	}

	log.Println("RUNNING APP")
	log.Println("RUNNING APP")
	log.Println("RUNNING APP")

	log.Println("SUCCESS")
}
