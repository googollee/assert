package assert

import (
	"strconv"
	"testing"

	"github.com/googollee/assert"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		input   string
		want    assert.Condition[int64]
		wantErr assert.Condition[error]
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
			if !tc.wantErr.Constrain(err) {
				t.Fatalf("%v.Apply(%v) fails", tc.wantErr, err)
			}
			if !tc.want.Constrain(got) {
				t.Errorf("%v.Apply(%v) fails", tc.want, tc.input)
			}
		})
	}
}
