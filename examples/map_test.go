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
				if l := len(v); l != 0 {
					t.Logf("len(keys) = %d, want: 0", l)
					return false
				}
				return true
			}, "empty"),
		},
		{
			name: "2Items",
			m: map[string]string{
				"a": "1",
				"b": "2",
			},
			wantKey: assert.AllOf(assert.Len[string](2), assert.Contain("a"), assert.Contain("b")),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var keys []string
			for key := range tc.m {
				keys = append(keys, key)
			}

			if should := tc.wantKey(keys); should != "" {
				t.Errorf("keys = %v, but should be %v", keys, should)
			}
		})
	}
}
