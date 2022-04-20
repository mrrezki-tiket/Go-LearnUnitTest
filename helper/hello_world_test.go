package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWord("Rezki")
	assert.Equalf(t, "Hello Rezki", result, "Result must be `Hello Rezki")
	fmt.Println("TestHelloWork with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWord("Nando ")
	assert.Equalf(t, "Hello Rezki", result, "Result must be `Hello Rezki")
	fmt.Println("TestHelloWork with assert done")
}

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
