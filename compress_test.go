package utils

import (
	"bytes"
	"testing"
)

func TestGzip(t *testing.T) {
	s := `232323`
	buffer := new(bytes.Buffer)
	_, err := Gzip([]byte(s), buffer)
	if err != nil {
		t.Fatal("gzip error", err)
	}

	output, err := Gunzip(buffer)
	if err != nil {
		t.Fatal("ungzip error:", err)
	}
	if string(output) != s {
		t.Fatal("ungzip unfinish:", string(output))
	}
}

func BenchmarkGzip(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := `fwefwefwe`
		buffer := new(bytes.Buffer)
		_, err := Gzip([]byte(s), buffer)
		if err != nil {
			b.Fatal("gzip error", err)
		}

		output, err := Gunzip(buffer)
		if err != nil {
			b.Fatal("ungzip error:", err)
		}
		if string(output) != s {
			b.Fatal("ungzip unfinish:", string(output))
		}
	}
}
