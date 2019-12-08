package spellchecker

import (
	"math"
	"sync"

	"github.com/spaolacci/murmur3"
)

type BloomFilter struct {
	ItemCount                uint32
	FalsePositiveProbability float64

	once      sync.Once
	size      uint32
	hashCount uint32
	bitMap    []bool
}

func (bf *BloomFilter) init() {
	bf.once.Do(func() {
		bf.size = bf.getSize()
		bf.hashCount = bf.getHashCount()
		bf.bitMap = make([]bool, bf.size)
	})
}

func (bf *BloomFilter) Add(s string) {
	bf.init()

	for i := uint32(0); i < bf.hashCount; i++ {
		d := murmur3.Sum32WithSeed([]byte(s), i) % bf.size
		bf.bitMap[d] = true
	}
}

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

// Return the size of bit array(m) to used using the following formula
//    m = -(n * lg(p)) / (lg(2)^2)
//    n : int - number of items expected to be stored in filter
//    p : float - False Positive probability in decimal
func (bf *BloomFilter) getSize() uint32 {
	return uint32(-(float64(bf.ItemCount) * math.Log(bf.FalsePositiveProbability)) / math.Pow(math.Log(2), 2))
}

// Return the hash function(k) to be used using following formula
//    k = (m/n) * lg(2)
//    m : int - size of bit array
//    n : int - number of items expected to be stored in filter
func (bf *BloomFilter) getHashCount() uint32 {
	return uint32((float64(bf.size) / float64(bf.ItemCount)) * math.Log(2))
}
