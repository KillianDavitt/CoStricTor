package main

import "testing"

func TestNewFilter(t *testing.T) {
	b := NewBloomFilter(5000, 8)
	data := []byte("Hi")
	b.Add(data, 0, 1)
}
