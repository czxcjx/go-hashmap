package hashmap

import (
	"math"
	"testing"
)

func TestCreate(t *testing.T) {
	h := New(10)
	if h.Size() != 0 {
		t.Error("Newly initialized hashmap doesn't have size 0")
	}
	if h.Capacity() != 10 {
		t.Error("Newly initialized hashmap has wrong capacity")
	}
}

func TestSetGet(t *testing.T) {
	h := New(3)
	if !h.Set("abc", 3) {
		t.Error("Set failed")
	}
	if h.Get("abc") != 3 {
		t.Error("Get failed")
	}
}

func TestUnsetValues(t *testing.T) {
	h := New(3)
	if !h.Set("abc", 3) {
		t.Error("Set failed")
	}
	if h.Get("abcd") != nil {
		t.Error("Get unset key did not return nil")
	}
	if h.Delete("abcd") != nil {
		t.Error("Delete unset key did not return nil")
	}
}

func TestSetDelete(t *testing.T) {
	h := New(3)
	if !h.Set("abc", 3) {
		t.Error("Set failed")
	}
	if h.Delete("abc") != 3 {
		t.Error("Delete failed")
	}
	if h.Get("abc") != nil {
		t.Error("Get deleted key did not return nil")
	}
	if !h.Set("abc", 5) {
		t.Error("Set failed")
	}
	if h.Get("abc") != 5 {
		t.Error("Get re-set key had wrong value")
	}
}

func TestSetReset(t *testing.T) {
	h := New(3)
	if !h.Set("abc", 3) {
		t.Error("Set failed")
	}
	if h.Size() != 1 {
		t.Error("Size was wrong")
	}
	if !h.Set("abc", 5) {
		t.Error("Set existing key failed")
	}
	if h.Size() != 1 {
		t.Error("Size was wrong")
	}
	if h.Get("abc") != 5 {
		t.Error("Get re-set key had wrong value")
	}
	if h.Delete("abc") != 5 {
		t.Error("Delete re-set key had wrong value")
	}
	if h.Size() != 0 {
		t.Error("Size was wrong")
	}
	if h.Get("abc") != nil {
		t.Error("Get re-set deleted key was not nil")
	}
}

func TestSetFullHashmap(t *testing.T) {
	h := New(3)
	if !h.Set("abc", 3) {
		t.Error("Set failed")
	}
	if !h.Set("def", 4) {
		t.Error("Set failed")
	}
	if !h.Set("ghi", 5) {
		t.Error("Set failed")
	}
	if h.Set("jkl", 6) {
		t.Error("Set succeeded on full hashmap")
	}
	if !h.Set("ghi", 6) {
		t.Error("Set failed on full hashmap but old key")
	}
	if h.Delete("ghi") != 6 {
		t.Error("Delete failed")
	}
	if !h.Set("jkl", 6) {
		t.Error("Set failed")
	}
	if h.Get("abc") != 3 {
		t.Error("Got wrong value")
	}
	if h.Get("def") != 4 {
		t.Error("Got wrong value")
	}
	if h.Get("jkl") != 6 {
		t.Error("Got wrong value")
	}
}

func TestLoadSizeCapacity(t *testing.T) {
	h := New(10)
	if !h.Set("abc", 5) {
		t.Error("Set failed")
	}
	if !h.Set("def", 5) {
		t.Error("Set failed")
	}
	if !h.Set("ghi", 5) {
		t.Error("Set failed")
	}
	if !h.Set("ghi", 7) {
		t.Error("Set failed")
	}
	if !float_equal(h.Load(), 3.0/10.0) {
		t.Error("Load was wrong")
	}
	if h.Size() != 3 {
		t.Error("Size was wrong")
	}
	if h.Capacity() != 10 {
		t.Error("Capacity was wrong")
	}
}

func float_equal(f1 float64, f2 float64) bool {
	EPSILON := 1e-9
	return math.Abs(f1-f2) < EPSILON
}
