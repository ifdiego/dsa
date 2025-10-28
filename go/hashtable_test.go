package hashtable

import "testing"

func TestHashTable(t *testing.T) {
	hashTable := CreateHashTable(10)
	hashTable.Insert("aabb")
	hashTable.Insert("bbcc")
	hashTable.Insert("abcd")

	tests := []struct {
		input    string
		expected string
		found    bool
	}{
		{"aabb", "aabb", true},
		{"bbcc", "bbcc", true},
		{"abcd", "abcd", true},
		{"xxxx", "", false},
	}

	for _, test := range tests {
		result := hashTable.Search(test.input)
		if test.found && (result == nil || *result != test.expected) {
			t.Errorf("search(%q) = %v, expected %q", test.input, result, test.expected)
		}
		if !test.found && result != nil {
			t.Errorf("search(%q) = %v, expected nil", test.input, *result)
		}
	}
}
