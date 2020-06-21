package suffixarray

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/google/go-cmp/cmp"
)

func TestManber_msd(t *testing.T) {
	tests := []struct {
		name string
		text string
		want *Manber
	}{
		{
			name: "normal",
			text: "abracadabra",
			want: &Manber{
				N:     11,
				Text:  "abracadabra",
				Index: []int{0, 3, 5, 7, 10, 1, 8, 4, 6, 2, 9, 11 /*sentinels*/},
				Rank:  []int{0, 5, 9, 0, 7, 0, 8, 0, 5, 9, 0, -1 /*sentinels*/},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.text)
			m.msd()
			opt := cmpopts.IgnoreUnexported(Manber{})
			if diff := cmp.Diff(m, tt.want, opt); diff != "" {
				t.Errorf("m.msd() differs: (-got +want)\n%s", diff)
			}
		})
	}
}