package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmarking

func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Hanif",
			request: "Hanif",
		},
		{
			name:    "Al-Irsyad",
			request: "Al-Irsyad",
		},
		{
			name:    "Hanif-Al",
			request: "Hanif Al",
		},
		{
			name:    "Hanif-Al-Irsyad",
			request: "Hanif Al Irsyad",
		},
	}

	for _, bench := range benchmark {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(bench.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Hanif", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hanif")
		}
	})
	b.Run("Hanif Al", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Hanif Al")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hanif")
	}
}

// Unit Testing

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Hanif",
			request:  "Hanif",
			expected: "Hello, Hanif",
		},
		{
			name:     "Al-Irsyad",
			request:  "Al-Irsyad",
			expected: "Hello, Al-Irsyad",
		},
		{
			name:     "Hanif-Al",
			request:  "Hanif Al",
			expected: "Hello, Hanif Al",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Hanif", func(t *testing.T) {
		result := HelloWorld("Hanif")
		require.Equal(t, "Hello, Hanif", result, "Result must be Hello, Hanif")
	})
	t.Run("Al-Irsyad", func(t *testing.T) {
		result := HelloWorld("Al Irsyad")
		require.Equal(t, "Hello, Al Irsyad", result, "Result must be Hello, Al Irsyad")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("Before Unit Test")

	m.Run()

	// after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can't run on Mac OS")
	}

	result := HelloWorld("Hanif")
	assert.Equal(t, "Hello, Hanif", result, "Result must be Hello, Hanif")

}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Hanif")
	assert.Equal(t, "Hello, Hanif", result, "Result must be Hello, Hanif")
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Hanif")
	require.Equal(t, "Hello, Hanif", result, "Result must be Hello, Hanif")
	fmt.Println("TestHelloWorld Done")
}
