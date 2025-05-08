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
			keyCheckers: cc.Func(func(t testing.TB, v []string) {
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
		{
			name: "10Items",
			m: map[string]string{
				"a": "1",
				"b": "2",
				"c": "3",
				"d": "4",
				"e": "5",
				"f": "6",
				"g": "7",
				"h": "8",
				"i": "9",
				"j": "10",
			},
			keyCheckers: cc.Equal([]string{"a"}),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			keys := slices.Collect(maps.Keys(tc.m))

			tc.keyCheckers.Want(t, keys)
		})
		// Output:
		// $ go test ./...
		// --- FAIL: TestMapRangeKeys (0.00s)
		// --- FAIL: TestMapRangeKeys/empty (0.00s)
		// map_test.go:22: got len(keys) = 0, want: 1
		// --- FAIL: TestMapRangeKeys/2Items (0.00s)
		// map_test.go:56: Defined at map_test.go:32: got len([a b])=2, want: 3
		// map_test.go:56: Defined at map_test.go:32: got [a b], want contain c
		// --- FAIL: TestMapRangeKeys/10Items (0.00s)
		// map_test.go:56: Defined at map_test.go:48: diff (-got, +want):
		// []string{
		// -   "d",
		// -   "e",
		// -   "f",
		// -   "g",
		// -   "h",
		// "a",
		// -   "b",
		// -   "c",
		// -   "i",
		// -   "j",
		// }
		// FAIL
		// FAIL    github.com/googollee/go-cc/examples     0.003s
		// FAIL
	}
}
