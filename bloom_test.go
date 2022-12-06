package main

import "testing"

func TestNewFilter(t *testing.T) {
	h:=20
	b := NewBloomFilter(5000, h)
	data := []byte("Hi")
	b.Add(data, 0, 1)
	t.Errorf("hi")

	d = b.data
	sumBits := 0
	for i:=0; i<len(d); i++ {
		sumBits += d[i]
	}
	if sumBits != h {
		t.Errorf("Perturbation happening for q=1, p=0")
	}
}
