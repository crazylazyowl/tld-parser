package tldparser

import (
	"testing"
)

func Test_TLD(t *testing.T) {
	p, _ := FromFile("tlds.dat")

	tests := []struct {
		hostname, tld string
	}{
		{"google.co.uk", "uk"},
		{"google.com", "com"},

		{"site.tld", ""},
	}

	for _, test := range tests {
		tld := p.TLD(test.hostname)
		if tld != test.tld {
			t.Fatalf("hostname = %s, tld = %s, answer = %s\n", test.hostname, test.tld, tld)
		}
	}
}

func benchmarkTLD(i int, b *testing.B) {
	p, _ := FromFile("tlds.dat")
	for n := 0; n < b.N; n++ {
		for x := 0; x < i; x++ {
			p.TLD("google.co.uk")
		}
	}
}

func Benchmark_TLD_10(b *testing.B) { benchmarkTLD(10, b) }

func Benchmark_TLD_100(b *testing.B) { benchmarkTLD(100, b) }

func Benchmark_TLD_1000(b *testing.B) { benchmarkTLD(1000, b) }

func Benchmark_TLD_100000(b *testing.B) { benchmarkTLD(100000, b) }

func Benchmark_TLD_1000000(b *testing.B) { benchmarkTLD(1000000, b) }
