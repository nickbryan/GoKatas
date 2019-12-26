package spellchecker

import (
	"testing"
)

type (
	adds   []string
	checks map[string]bool
)

func TestBloomFilter(t *testing.T) {
	var tests = map[string]struct {
		adds   adds
		checks checks
	}{
		"false for not in empty filter":              {adds: adds{}, checks: checks{"test": false}},
		"true for single item":                       {adds: adds{"test"}, checks: checks{"test": true}},
		"false for single item not exists":           {adds: adds{"test"}, checks: checks{"tree": false}},
		"true for both items in set":                 {adds: adds{"test", "test2"}, checks: checks{"test": true, "test2": true}},
		"false for one items not in two item in set": {adds: adds{"test", "test2"}, checks: checks{"test": true, "test3": false}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			bf := &BloomFilter{
				ItemCount:                uint32(len(tc.adds)),
				FalsePositiveProbability: 0.05,
			}
			for _, a := range tc.adds {
				bf.Add(a)
			}
			for v, want := range tc.checks {
				if got := bf.Exists(v); got != want {
					t.Errorf("expected Exists to return [%t] for vale [%s] got: [%t]", want, v, got)
				}
			}
		})
	}
}
