package main

import "testing"
import "fmt"

func TestNewCrews(t *testing.T) {
	// filter size, hashes, num_websites, pthresh, sthresh, p, q
	c := NewCrews(1000, 8, 1, 1, 1, 0, 1)
	c.ReportHsts("www.google.com")
	if 8 != sumFilter(c.primary) {
		t.Errorf("Error in reporting")
	}

	c.ReportHttp("www.example.com")
	if 8 != sumFilter(c.secondary) {
		t.Errorf("Error in reporting")
	}

	if !(1 >= uint(float64(( c.primary.count/c.numWebsites))*c.primaryThresholdModifier)){
		t.Errorf("Not going to get a good result from primary")
	}

	count := c.primary.Test([]byte("www.google.com"))
	if count != 1 {
		fmt.Printf("Count: %d\n",count)
		fmt.Println(c.primary.data)
		t.Errorf("bad initial count")
	}
	adjustedCount := uint((float64(count) - c.p * float64(c.primary.count))/(c.q-c.p))
	if adjustedCount != 1 {
		t.Errorf("Bad adjusted count")
	}
	
	if !c.PrimaryTest("www.google.com") {
		fmt.Println( c.primary.count/c.numWebsites)
		t.Errorf("Getting wrong result from primaryTest")
	}

	if !c.SecondaryTest("www.example.com") {
		fmt.Println( c.secondary.count/c.numWebsites)
		t.Errorf("Getting wrong result from secondaryTest")
	}

	if c.PrimaryTest("www.example.com") {
		t.Errorf("Getting wrong result from primaryTest")
	}

	if c.SecondaryTest("www.google.com") {
		t.Errorf("Getting wrong result from secondaryTest")
	}
	
}

func TestManyInsertionsCrews(t *testing.T) {
	c := NewCrews(1000, 8, 1, 1, 1, 0, 1)
	for i:=0; i<5; i++ {
		c.ReportHsts("www.google.com")
	}

	if !c.PrimaryTest("www.google.com"){
		t.Errorf("Fail on multiple insertions")
	}

	if c.PrimaryTest("www.example.com"){
		t.Errorf("Fail multiple insertions")
	}
}


func sumFilter(b *BloomFilter) (uint) {
	var sum uint = 0
	for i:= 0; i<len(b.data); i++ {
		sum += b.data[i]
	}
	return sum
}
