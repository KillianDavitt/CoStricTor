package main

import "testing"

func TestNewCrews(t *testing.T) {
	c := NewCrews(1000, 8, 50, 0.00000001, 0.01, 0, 1)
	c.ReportHsts("www.google.com")
	if 8 != sumFilter(c.primary) {
		t.Errorf("Error in reporting")
	}

	c.ReportHttp("www.example.com")
	if 8 != sumFilter(c.secondary) {
		t.Errorf("Error in reporting")
	}

	if c.PrimaryTest("www.google.com") {
		fmt.Println( c.primary.count/c.numWebsites)
		t.Errorf("Getting wrong result from primaryTest")
	}
}

func sumFilter(b *BloomFilter) (uint) {
	var sum uint = 0
	for i:= 0; i<len(b.data); i++ {
		sum += b.data[i]
	}
	return sum
}
