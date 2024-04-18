package lrucache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	// ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	// [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// [null, null, null, 1, null, -1, null, -1, 3, 4]

	cache := New(2)

	// Test case 1: Put key 1 with value 1
	cache.Put(1, 1)

	// Test case 2: Put key 2 with value 2
	cache.Put(2, 2)

	// Test case 3: Get value for key 1, expected output: 1
	if value := cache.Get(1); value != 1 {
		t.Errorf("Expected value for key 1: 1, but got %d", value)
	}

	// Test case 4: Put key 3 with value 3, which will evict key 2
	cache.Put(3, 3)

	// Test case 5: Get value for key 2, expected output: -1 (evicted)
	if value := cache.Get(2); value != -1 {
		t.Errorf("Expected value for key 2: -1, but got %d", value)
	}

	// Test case 6: Put key 4 with value 4, which will evict key 1
	cache.Put(4, 4)

	// Test case 7: Get value for key 1, expected output: -1 (evicted)
	if value := cache.Get(1); value != -1 {
		t.Errorf("Expected value for key 1: -1, but got %d", value)
	}

	// Test case 8: Get value for key 3, expected output: 3
	if value := cache.Get(3); value != 3 {
		t.Errorf("Expected value for key 3: 3, but got %d", value)
	}

	// Test case 9: Get value for key 4, expected output: 4
	if value := cache.Get(4); value != 4 {
		t.Errorf("Expected value for key 4: 4, but got %d", value)
	}
}
