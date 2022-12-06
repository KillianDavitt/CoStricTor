package main

import "testing"

func TestNewFilter(t *testing.T) {
	var h uint =20
	b := NewBloomFilter(5000, h)
	data := []byte("Hi")
	b.Add(data, 0, 1)

	d := b.data
	var sumBits uint = 0
	for i:=0; i<len(d); i++ {
		sumBits += d[i]
	}
	if sumBits != h {
		t.Errorf("Perturbation happening for q=1, p=0")
	}
}
