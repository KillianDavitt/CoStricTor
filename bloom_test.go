package main

import "testing"

func TestNewFilter(t *testing.T) {
	hs := [5]uint{1,4,6,18,22}
	for i:=0; i<len(hs); i++ {
		b := NewBloomFilter(100, hs[i])
		data := []byte("Hi")
		b.Add(data, 0, 1)

		d := b.data
		var sumBits uint = 0
		for j:=0; j<len(d); j++ {
			sumBits += d[j]
		}
		if sumBits != hs[i] {
			t.Errorf("Perturbation happening for q=1, p=0")
		}
	}
}
