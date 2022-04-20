package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWord("Rezki")
	if result != "Hello Rezki" {
		// unit test failed
		panic("Result is not Hello Rezki ")
	}
}
