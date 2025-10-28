package main

import (
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkTrieLookup(b *testing.B) {
	t := CreateTrie(5)
	rand.Seed(42)
	numEntries := 10000
	keys := make([]string, numEntries)
	for i := 0; i < numEntries; i++ {
		key := "key" + strconv.Itoa(i) + string(rune('a'+rand.Intn(26)))
		keys[i] = key
		t.Insert(key)
	}

	testKey := keys[len(keys)-1]

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = t.Search(testKey)
	}
}
