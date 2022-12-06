package main

import "testing"

func TestNewCrews(t *testing.T) {
	c := NewCrews(1000, 8, 50, 0.01, 0.01, 0, 1)
	c.ReportHsts("www.google.com")
	if 8 != sumFilter(c.primary) {
		t.Errorf("Error in reporting")
	}
}

func sumFilter(b *BloomFilter) (int) {
	var sum uint = 0
	for i:= 0; i<len(b.data); i++ {
		sum += b.data[i]
	}
	return sum
}
