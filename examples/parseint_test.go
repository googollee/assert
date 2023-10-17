package assert

import (
	"strconv"
	"testing"

	"github.com/googollee/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		input   string
		want    assert.Asserter
		wantErr assert.Asserter
	}{
		{
			input:   "",
			want:    assert.Any(),
			wantErr: assert.ErrorContains("invalid syntax"),
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
			want:    assert.Any(),
			wantErr: assert.ErrorContains("out of range"),
		},
		{
			input:   "-9223372036854775808",
			want:    assert.Equal(-1 << 63),
			wantErr: assert.IsNil[error](),
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := strconv.ParseInt(tc.input, 10, 64)
			tc.wantErr.FailNow(err, "ParseInt(%d, 10, 64) error", tc.input)
			tc.want.Check(got, "ParseInt(%d, 10, 64)", tc.input)
		})
	}
}
