package main

import "testing"

func TestNewCrews(t *testing.T) {
	c := NewCrews(1000, 8, 50, 0.01, 0.01, 0, 1)
	c.ReportHsts("www.google.com")
}
