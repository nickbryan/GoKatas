package spellchecker

import (
	"math"
	"sync"

	"github.com/spaolacci/murmur3"
)

// BloomFilter will let us know if a word is either definitely is not in the set or possibly be in the set.
type BloomFilter struct {
	// ItemCount should be set to the expected number of elements that we will store in the filter.
	ItemCount uint32

	// FalsePositiveProbability should indicate the probability that we are comfortable with encountering a false positive.
	// It should be set to a number been 0 and 1.
	FalsePositiveProbability float64

	once      sync.Once
	size      uint32
	hashCount uint32
	bitMap    []bool
}

// Add sets the relevant bits in the filter to true for the given string.
func (bf *BloomFilter) Add(s string) {
	bf.init()

	for i := uint32(0); i < bf.hashCount; i++ {
		d := murmur3.Sum32WithSeed([]byte(s), i) % bf.size
		bf.bitMap[d] = true
	}
}

// Exists checks the relevant bits in the filter, if any of them are set to false then we know that the given string
// is definitely not in the set.
func (bf *BloomFilter) Exists(s string) bool {
	bf.init()

	if len(bf.bitMap) == 0 {
		return false
	}

	for i := uint32(0); i < bf.hashCount; i++ {
		d := murmur3.Sum32WithSeed([]byte(s), i) % bf.size
		if bf.bitMap[d] == false {
			return false
		}
	}

	return true
}

// init can be called lazily to initialise the filter. It will only run the initialisation code once.
func (bf *BloomFilter) init() {
	bf.once.Do(func() {
		bf.size = bf.calculateSize()
		bf.hashCount = bf.calculateHashCount()
		bf.bitMap = make([]bool, bf.size)
	})
}

// calculateSize returns the size of the bitmap to be used using the following formula:
//    m = -(n * lg(p)) / (lg(2)^2)
//    n : int - number of items expected to be stored in filter
//    p : float - False Positive probability in decimal
func (bf *BloomFilter) calculateSize() uint32 {
	return uint32(-(float64(bf.ItemCount) * math.Log(bf.FalsePositiveProbability)) / math.Pow(math.Log(2), 2))
}

// calculateHashCount returns the number of hash functions to be used using following formula:
//    k = (m/n) * lg(2)
//    m : int - size of bit array
//    n : int - number of items expected to be stored in filter
func (bf *BloomFilter) calculateHashCount() uint32 {
	return uint32((float64(bf.size) / float64(bf.ItemCount)) * math.Log(2))
}
