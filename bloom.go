package main

import (
	"hash"
	"hash/fnv"
	//"math/rand"
	//"fmt"
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
			count: 0,
	}
}

func (b *BloomFilter) Add(data []byte, p float64, q float64) *BloomFilter {
	lower, upper := hashKernel(data, b.hash)
	adq := uint32(q * float64(4294967295.0))
	adp := uint32(p * float64(4294967295.0))
	//adr := uint32(2147483647)
	var newData []uint = make([]uint, b.filterSize)
	for i := uint(0); i < b.numHashes; i++ {
		trueBit := ((lower+upper*i)%b.filterSize)
		newData[trueBit]+=1

	}
	falseBits := 0
	for i:=uint(0); i<b.filterSize; i++ {
		r := fastrand.FastRand()
		if newData[i]==1 {
			if r>=adq {
				newData[i]=0
			}
		} else {
			if r<adp {
				newData[i]=1
				falseBits+=1
			}
		}
		b.data[i]+=newData[i]
	}
	fmt.Println(falseBits)
	b.count++
	return b
}

func (b *BloomFilter) Test(data []byte) uint {
	lower, upper := hashKernel(data, b.hash)
	var result []uint = make([]uint, b.numHashes);
	// Get the bit counts for each hash function
	for i := uint(0); i < b.numHashes; i++ {
		trueBit := ((lower+upper*i)%b.filterSize)		
		result[i] = b.data[trueBit]
	}
	var min uint = 0
	//fmt.Println(result)
	for i, e := range result {
		if i==0 || e < min {
			min = e
		}
	}
	return min;
}

func hashKernel(data []byte, hash hash.Hash64) (uint, uint) {
	hash.Write(data)
	sum := hash.Sum64()
	hash.Reset()
	// Separating the bits out seems odd, but is useful for filter indexing
	upper := uint(sum & 0xffffffff)
	lower := uint((sum >> 32) & 0xffffffff)
	return upper, lower
}
