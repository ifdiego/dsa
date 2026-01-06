package main

import "testing"

func TestHashTable(t *testing.T) {
	ht := NewHashTable(10)

	ht.Insert("name", "John")
	ht.Insert("age", "99")
	ht.Insert("country", "Brazil")

	if name, ok := ht.Search("name"); !ok || name != "John" {
		t.Errorf("Expected John, but got %v", name)
	}
	if age, ok := ht.Search("age"); !ok || age != "99" {
		t.Errorf("Expected 99, but got %v", age)
	}

	ht.Remove("country")
	if country, ok := ht.Search("country"); ok {
		t.Errorf("Expected Brazil to be removed, but got %v", country)
	}

	if _, ok := ht.Search("notfound"); ok {
		t.Errorf("Expected a key to not be found")
	}
}
