package main

import (
	"hash"
	"hash/fnv"
	//"math/rand"
	"github.com/detailyang/fastrand-go"
)

// The design of this structure is adapted from https://github.com/tylertreat/BoomFilters
type BloomFilter struct {
	data []uint    // filter data
	hash    hash.Hash64 
	filterSize       uint        
	numHashes       uint       
	count   uint      
}

func NewBloomFilter(filterSize uint, numHashes uint) *BloomFilter {
		return &BloomFilter{
		data: make([]uint, filterSize),
		hash:    fnv.New64(),
		filterSize:       filterSize,
		numHashes:       numHashes,
	}
}

func (b *BloomFilter) Add(data []byte, p float64, q float64) *BloomFilter {
	lower, upper := hashKernel(data, b.hash)
	adq := uint32(q * float64(4294967295.0))
	adp := uint32(p * float64(4294967295.0))
	adr := uint32(2147483647)
	for i := uint32(0); i < uint32(b.numHashes); i++ {
		trueBit := ((uint32(lower)+uint32(upper)*i)%b.filterSize)
		for j:= uint32(0); j<b.filterSize; j++ {
			var r uint32;
			if q==1 && p==0 {
				// 0.5 is an arbritrary number which is less than 1 and greater than 0
				r = adr
			} else {
				r = fastrand.FastRand()
			}
			if j==trueBit {
				// q chance of returning 1
				if r<adq {
					b.data[j]+=1
				}
			} else {
				// p chance of returning 1
				if r<adp {
					b.data[j] += 1
				}
			}
		}
	}

	b.count++
	return b
}

func (b *BloomFilter) Test(data []byte) uint {
	lower, upper := hashKernel(data, b.hash)
	var result []uint = make([]uint, b.numHashes);
	// Get the bit counts for each hash function
	for i := uint(0); i < b.numHashes; i++ {
		result[i] = b.data[((uint(lower)+uint(upper)*i)%b.filterSize)]
	}
	var min uint;
	for i, e := range result {
		if i==0 || e < min {
			min = e
		}
	}
	return min;
}

func hashKernel(data []byte, hash hash.Hash64) (uint32, uint32) {
	hash.Write(data)
	sum := hash.Sum64()
	hash.Reset()
	// Separating the bits out seems odd, but is useful for filter indexing
	upper := uint32(sum & 0xffffffff)
	lower := uint32((sum >> 32) & 0xffffffff)
	return upper, lower
}
