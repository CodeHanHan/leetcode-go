package bloomfilter

import (
	"github.com/spaolacci/murmur3"
)

var (
	defaultMaps int = 10
	maxMaps     int = 20
	defaultLens int = 20
	maxLens     int = 50
)

type Hash func([]byte) uint32

type BloomFilter struct {
	// 哈希函数数量
	maps int
	// 哈希函数
	hashFn Hash
	// 位点存储
	store []uint8
}

func NewBloomFilter(maps int, hashFn Hash, storeLens int) *BloomFilter {
	if maps > maxMaps {
		maps = maxMaps
	}

	if storeLens > maxLens {
		storeLens = maxLens
	}

	if hashFn == nil {
		hashFn = murmur3.Sum32
	}

	return &BloomFilter{
		maps:   maps,
		hashFn: hashFn,
		store:  make([]uint8, storeLens),
	}
}

func (b *BloomFilter) Add(v []byte) {
	for i := 0; i < b.maps; i++ {
		idx := b.hashFn(append(v, byte(i))) % (uint32(len(b.store)) * 8)
		to1(b.store, idx)
	}
}

func (b *BloomFilter) Exist(v []byte) bool {
	for i := 0; i < b.maps; i++ {
		idx := b.hashFn(append(v, byte(i))) % (uint32(len(b.store)) * 8)
		if !is1(b.store, idx) {
			return false
		}
	}

	return true
}

func to1(store []uint8, idx uint32) {
	i, j := idx/8, idx%8
	store[i] |= 1 << j
}

func is1(store []uint8, idx uint32) bool {
	i, j := idx/8, idx%8
	return (store[i] >> j & 0x1) == 1
}
