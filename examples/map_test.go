package assert

import (
	"maps"
	"slices"
	"testing"

	"github.com/googollee/assert"
)

func TestMapRangeKeys(t *testing.T) {
	tests := []struct {
		name        string
		m           map[string]string
		keyCheckers assert.Assert[[]string]
	}{
		{
			name: "empty",
			m:    map[string]string{},
			keyCheckers: assert.Func(func(t assert.T, v []string) {
				if l := len(v); l != 1 {
					t.Checkf("got len(keys) = %d, want: 1", l)
				}
			}),
		},
		{
			name: "2Items",
			m: map[string]string{
				"a": "1",
				"b": "2",
			},
			keyCheckers: assert.AllOf(assert.Len[string](3), assert.Contain("a"), assert.Contain("b"), assert.Contain("c")),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keys := slices.Collect(maps.Keys(tc.m))

			tc.keyCheckers.Want(t, keys)
			// Output:
			// $ go test ./...
			// ?       github.com/googollee/assert     [no test files]
			// --- FAIL: TestMapRangeKeys (0.00s)
			//     --- FAIL: TestMapRangeKeys/empty (0.00s)
			//         map_test.go:22: got len(keys) = 0, want: 1
			//     --- FAIL: TestMapRangeKeys/2Items (0.00s)
			//         map_test.go:40: got len([a b])=2, want: 3
			//         map_test.go:40: got [a b], want contain c
			// FAIL
			// FAIL    github.com/googollee/assert/examples    0.001s
			// FAIL
		})
	}
}
