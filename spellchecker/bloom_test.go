package spellchecker

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	var tests = map[string]struct {
		additions []string
		tests     map[string]bool
	}{
		"false for not in empty filter":              {additions: []string{}, tests: map[string]bool{"test": false}},
		"true for single item":                       {additions: []string{"test"}, tests: map[string]bool{"test": true}},
		"false for single item not exists":           {additions: []string{"test"}, tests: map[string]bool{"tree": false}},
		"true for both items in set":                 {additions: []string{"test", "test2"}, tests: map[string]bool{"test": true, "test2": true}},
		"false for one items not in two item in set": {additions: []string{"test", "test2"}, tests: map[string]bool{"test": true, "test3": false}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bf := &BloomFilter{
				ItemCount:                uint32(len(tc.additions)),
				FalsePositiveProbability: 0.05,
			}
			for _, a := range tc.additions {
				bf.Add(a)
			}
			for v, want := range tc.tests {
				if got := bf.Exists(v); got != want {
					t.Errorf("expected Exists to return [%t] for vale [%s] got: [%t]", want, v, got)
				}
			}
		})
	}
}
