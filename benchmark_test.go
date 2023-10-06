package main

import "testing"

func BenchmarkCSVtoJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CSVtoJSON()
	}
}
