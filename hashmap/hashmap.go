// Package hashmap provides a fixed-size hash table with string keys and 
// generic values using linear probing to resolve collisions
package hashmap

import "math/rand"

type entry struct {
	isFilled bool
	key      string
	value    interface{}
}

type HashMap struct {
	entries []entry
	size    int

	// Universal hash family (Ax + B) % P
	// Using a universal hash family ensures that collisions (and near-collisions)
  // are rarer. Since we are using linear probing, having many strings hash to 
  // nearby buckets causes the length of a "chain" of filled buckets to be long,
  // which makes operations slow. When testing this compared to just using 
  // x % P, with strings "1", "2", etc., over 10x speedup was observed.
	hashA uint64
	hashB uint64
	hashP uint64
}

// New allocates a new hash map with the given size
func New(size int) *HashMap {
	h := new(HashMap)
	h.entries = make([]entry, size)
	h.size = 0
	// Initialize P to a large prime
	h.hashP = 5112733757
  // Initialize A and B to random numbers in [1, P) and [0, P)
	h.hashA = uint64(rand.Int63n(int64(h.hashP)-1) + 1)
	h.hashB = uint64(rand.Int63n(int64(h.hashP)))
	return h
}

// Capacity returns the maximum capacity of the hashmap
func (h *HashMap) Capacity() int {
	return len(h.entries)
}

// Size returns the number of elements currently in the hashmap
func (h *HashMap) Size() int {
	return h.size
}

// Load returns the load factor of the hashmap (i.e. 
// number of elements / capacity). Since it is a fixed-size
// hashmap this will always be less than or equal to 1
func (h *HashMap) Load() float64 {
	return float64(h.Size()) / float64(h.Capacity())
}

// Set will insert the key-value pair (key, value) into the hashmap. If a key 
// already exists in the hashmap, Set will overwrite the corresponding value.
// If there is no capacity to store the new key, Set will fail.
// Set returns true if the insertion was successful, false otherwise.
func (h *HashMap) Set(key string, value interface{}) bool {
	bucket, _ := h.findEntryOrEmptySlot(key)
	if bucket == nil {
		// No more space
		return false
	} else if bucket.isFilled {
		// Existing element
		bucket.value = value
	} else {
		// New element
		*bucket = entry{isFilled: true, key: key, value: value}
		h.size++
	}
	return true
}

// Get will retrieve the value for the given key from the hashmap. If a
// key does not exist in the hashmap, Get will return nil.
func (h *HashMap) Get(key string) interface{} {
	bucket, _ := h.findEntryOrEmptySlot(key)
	if bucket == nil || !bucket.isFilled {
		// Looped through whole array or hit an empty bucket
		return nil
	} else {
		// Found the correct key
		return bucket.value
	}
}

// Delete will delete the given key from the hashmap and return it. If a 
// key does not exist in the hashmap, Delete will not do anything and return 
// nil.
func (h *HashMap) Delete(key string) interface{} {
	bucket, bucketIndex := h.findEntryOrEmptySlot(key)
	if bucket == nil || !bucket.isFilled {
		// Looped through whole array or hit an empty bucket
		return nil
	} else {
		// Found value, unset the entry in the table
		returnValue := bucket.value
		emptyEntry := entry{isFilled: false, key: "", value: nil}
		*bucket = emptyEntry
		h.size--

		// Rehash all later elements until we find an unfilled entry to bridge the gap
		bucketIndex = (bucketIndex + 1) % h.Capacity()
		for h.entries[bucketIndex].isFilled {
			entryKey := h.entries[bucketIndex].key
			entryValue := h.entries[bucketIndex].value
			h.entries[bucketIndex] = emptyEntry
			h.size--
			h.Set(entryKey, entryValue)
			bucketIndex = (bucketIndex + 1) % h.Capacity()
		}
		return returnValue
	}
}

// Hash function f(x) = ((A * x + B) % P) % capacity
// Converts a string into a base-256 number (mod P)
func (h *HashMap) hash(input string) int {
	var sum uint64 = 0
	// Treat string as a base-256 number (mod P)
	for _, char := range []byte(input) {
		sum = (sum*256 + uint64(char)) % h.hashP
	}
	// Calculate A * x + B (mod P)
	sum = (h.hashB + sum*h.hashA) % h.hashP
	return int(sum % uint64(h.Capacity()))
}

// Finds the entry with a given key, or an empty slot at the end of that block
func (h *HashMap) findEntryOrEmptySlot(key string) (*entry, int) {
	// Get directly hashed slot
	bucketIndex := h.hash(key)

	// Loop through indices until we find an empty slot or a slot that matches the key
	for i := 0; i < h.Capacity(); i++ {
		if h.entries[bucketIndex].isFilled {
			if h.entries[bucketIndex].key == key {
				return &h.entries[bucketIndex], bucketIndex
			}
		} else {
			return &h.entries[bucketIndex], bucketIndex
		}

		// Increment and loop-around the index
		bucketIndex++
		if bucketIndex == h.Capacity() {
			bucketIndex = 0
		}
	}
	return nil, -1
}

