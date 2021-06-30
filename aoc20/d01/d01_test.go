package d01_test

import (
	"io"
	"strings"
	"testing"

	. "github.com/clfs/x/aoc20/d01"
	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	t.Parallel()
	cases := map[string]struct {
		r    io.Reader
		want []int
	}{
		"length 1": {
			r:    strings.NewReader("1"),
			want: []int{1},
		},
		"length 2": {
			r:    strings.NewReader("1\n2"),
			want: []int{1, 2},
		},
	}
	for name, tc := range cases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := Parse(tc.r)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Parse() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
