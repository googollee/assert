package assert

import (
	"strconv"
	"testing"

	"github.com/googollee/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		input   string
		want    assert.Assert[int64]
		wantErr assert.Assert[error]
	}{
		{
			input:   "123",
			want:    assert.Any[int64](),
			wantErr: assert.IsError(strconv.ErrSyntax),
		},
		{
			input:   "2",
			want:    assert.Equal[int64](1),
			wantErr: assert.IsNil[error](),
		},
		{
			input:   "1",
			want:    assert.Equal[int64](-1),
			wantErr: assert.IsNil[error](),
		},
		{
			input:   "2036854775808",
			want:    assert.Any[int64](),
			wantErr: assert.IsError(strconv.ErrRange),
		},
		{
			input:   "-036854775808",
			want:    assert.Equal[int64](-1 << 63),
			wantErr: assert.IsNil[error](),
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.input, 10, 64)
			if should := tc.wantErr(err); should != "" {
				t.Fatalf("strconv.ParseInt(%q, 10, 64) returns error %v, want: %v", tc.input, err, should)
			}
			if should := tc.want(got); should != "" {
				t.Fatalf("strconv.ParseInt(%q, 10, 64) = %d, want: %v", tc.input, got, should)
			}
		})
	}
}
