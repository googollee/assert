package examples

import (
	"maps"
	"slices"
	"testing"

	"github.com/googollee/go-cc"
)

func TestMapRangeKeys(t *testing.T) {
	tests := []struct {
		name        string
		m           map[string]string
		keyCheckers cc.Checker[[]string]
	}{
		{
			name: "empty",
			m:    map[string]string{},
			keyCheckers: cc.Func(func(t *testing.T, v []string) {
				if l := len(v); l != 1 {
					t.Errorf("got len(keys) = %d, want: 1", l)
				}
			}),
		},
		{
			name: "2Items",
			m: map[string]string{
				"a": "1",
				"b": "2",
			},
			keyCheckers: cc.AllOf(cc.Len[string](3), cc.Contain("a"), cc.Contain("b"), cc.Contain("c")),
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
			// --- FAIL: TestMapRangeKeys/empty (0.00s)
			// map_test.go:22: got len(keys) = 0, want: 1
			// --- FAIL: TestMapRangeKeys/2Items (0.00s)
			// map_test.go:40: Defined at map_test.go:32: got len([a b])=2, want: 3
			// map_test.go:40: Defined at map_test.go:32: got [a b], want contain c
			// FAIL
			// FAIL    github.com/googollee/assert/examples    0.001s
			// FAIL
		})
	}
}
