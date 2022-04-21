package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("Rezki")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		Expected string
	}{
		{
			Name:     "Rezki",
			Request:  "Rezki",
			Expected: "Hello Rezki",
		},
		{
			Name:     "Nando",
			Request:  "Nando",
			Expected: "Hello Nando",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := HelloWord(test.Request)
			require.Equal(t, test.Expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Rezki", func(t *testing.T) {
		result := HelloWord("Rezki")
		require.Equal(t, "Hello Rezki", result)
	})
	t.Run("Nando", func(t *testing.T) {
		result := HelloWord("Nando ")
		require.Equal(t, "Hello Nando", result)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run in Mac OS")
	}

	TestHelloWorldAssert(t)

}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWord("Rezki")
	assert.Equalf(t, "Hello Rezki", result, "Result must be `Hello Rezki")
	fmt.Println("TestHelloWork with assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWord("Nando ")
	require.Equalf(t, "Hello Rezki", result, "Result must be `Hello Rezki")
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
