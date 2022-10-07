package main

import (
	"hash"
	"hash/fnv"
)

// The design of this structure is adapted from https://github.com/tylertreat/BoomFilters
type BloomFilter struct {
	data []uint    // filter data
	hash    hash.Hash64 // hash function (kernel for all k functions)
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

func (b *BloomFilter) Add(data []byte) *BloomFilter {
	lower, upper := hashKernel(data, b.hash)

	for i := uint(0); i < b.numHashes; i++ {
		// Times the upper bit by i to ensure a different index per hash function
		b.data[((uint(lower)+uint(upper)*i)%b.filterSize)]+=1
	}

	b.count++
	return b
}

func (b *BloomFilter) Test(data []byte) []uint {
	lower, upper := hashKernel(data, b.hash)
	var result []uint = make([]uint, b.numHashes);
	// Get the bit counts for each hash function
	for i := uint(0); i < b.numHashes; i++ {
		result[i] = b.data[((uint(lower)+uint(upper)*i)%b.filterSize)]
	}

	return result;
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
