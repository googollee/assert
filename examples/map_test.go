package assert

import (
	"testing"

	"github.com/googollee/assert"
)

func TestMapRangeKeys(t *testing.T) {
	tests := []struct {
		name    string
		m       map[string]string
		wantKey assert.Assert[[]string]
	}{
		{
			name: "empty",
			m:    map[string]string{},
			wantKey: assert.Func(func(v []string) bool {
				if l := len(v); l != 1 {
					t.Logf("len(keys) = %d, want: 1", l)
					return false
				}
				return true
			}, "has 1 entry"),
		},
		{
			name: "2Items",
			m: map[string]string{
				"a": "1",
				"b": "2",
			},
			wantKey: assert.AllOf(assert.Len[string](3), assert.Contain("a"), assert.Contain("b"), assert.Contain("c")),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var keys []string
			for key := range tc.m {
				keys = append(keys, key)
			}

			if should := tc.wantKey(keys); should != "" {
				t.Errorf("keys = %v, want: %v", keys, should)
			}
		})
	}
}
