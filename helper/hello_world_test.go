package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWord("Nando")
	if result != "Hello Rezki" {
		// unit test failed
		t.Error("Result must be `Hello Rezki")
	}

	fmt.Println("TestHelloWord done")
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWord("Ekki")
	if result != "Hello Rezki" {
		// unit test failed
		t.Fatal("Result must be `Hello Rezki")
	}

	fmt.Println("TestHelloWord done")
}
