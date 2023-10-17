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
			name:    "empty",
			m:       map[string]string{},
			wantKey: assert.Func(func(v []string) bool { return len(v) == 0 }),
		},
		{
			name: "2Items",
			m: map[string]string{
				"a": "1",
				"b": "2",
			},
			wantKey: assert.All(assert.Len[string](2), assert.Contain("a"), assert.Contain("b")),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var keys []string
			for key := range tc.m {
				keys = append(keys, key)
			}

			tc.wantKey.Check(t, keys)
		})
	}
}
