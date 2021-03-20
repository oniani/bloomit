package main

import (
	"github.com/twmb/murmur3"
	"github.com/willf/bitset"
	"math"
)

// A Bloom Filter data structure
type BloomFilter struct {
	m uint           // Number of bits in the bit vector
	k uint           // Number of hash functions
	b *bitset.BitSet // Bit set
}

// New creates a Bloom filter with optimal number of hash functions
func New(n uint, eps float64) *BloomFilter {
	// Optimal number of bits in the bit vector
	m := uint(math.Ceil(-1 * float64(n) * math.Log(eps) / math.Pow(math.Log(2), 2)))

	// Optimal number of hash functions
	k := uint(math.Ceil(float64(m) / float64(n) * math.Log(2)))

	// Define the bit vector
	b := bitset.New(m)

	// Create and return a Bloom Filter
	return &BloomFilter{m: m, k: k, b: b}
}

// Add an item to the bloom filter
func (bf *BloomFilter) Add(data string) {
	for i := uint(0); i < bf.k; i++ {
		hash := murmur3.SeedStringSum64(uint64(i), data)
		idx := uint(hash % uint64(bf.m))
		bf.b.Set(idx)
	}
}

// Check whether the given item is present. If the item is not present, it is
// definitely not present. If the item is present, however, there is a chance
// that it is a false positive.
func (bf *BloomFilter) Check(data string) bool {
	for i := uint(0); i < bf.k; i++ {
		hash := murmur3.SeedStringSum64(uint64(i), data)
		idx := uint(hash % uint64(bf.m))
		if !bf.b.Test(idx) {
			return false
		}
	}

	return true
}
