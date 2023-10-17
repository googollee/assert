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
			input:   "",
			want:    assert.Any[int64](),
			wantErr: assert.IsError(strconv.ErrSyntax),
		},
		{
			input:   "1",
			want:    assert.Equal[int64](1),
			wantErr: assert.IsNil[error](),
		},
		{
			input:   "-1",
			want:    assert.Equal[int64](-1),
			wantErr: assert.IsNil[error](),
		},
		{
			input:   "9223372036854775808",
			want:    assert.Any[int64](),
			wantErr: assert.IsError(strconv.ErrRange),
		},
		{
			input:   "-9223372036854775808",
			want:    assert.Equal[int64](-1 << 63),
			wantErr: assert.IsNil[error](),
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.input, 10, 64)
			tc.wantErr.Ensuref(t, err, "ParseInt(%d, 10, 64) error", tc.input)
			tc.want.Checkf(t, got, "ParseInt(%d, 10, 64)", tc.input)
		})
	}
}
