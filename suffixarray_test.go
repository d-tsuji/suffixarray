package suffixarray

import (
	"index/suffixarray"
	"io/ioutil"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBuild(t *testing.T) {
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
				Index: []int{10, 7, 0, 3, 5, 8, 1, 4, 6, 9, 2, 11 /*sentinels*/},
				Rank:  []int{2, 6, 10, 3, 7, 4, 8, 1, 5, 9, 0, -1 /*sentinels*/},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.text)
			m.Build()
			opt := cmpopts.IgnoreUnexported(Manber{})
			if diff := cmp.Diff(m, tt.want, opt); diff != "" {
				t.Errorf("m.Build() differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestLookupAll(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		target string
		want   []int
	}{
		{
			name:   "normal",
			text:   "banana",
			target: "ana",
			want:   []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := New(tt.text)
			m.Build()

			// This Transformer sorts a []int.
			trans := cmp.Transformer("Sort", func(in []int) []int {
				out := append([]int(nil), in...) // Copy input to avoid mutating it
				sort.Ints(out)
				return out
			})

			got := m.LookupAll(tt.target)
			if diff := cmp.Diff(got, tt.want, trans); diff != "" {
				t.Errorf("m.LookupAll() differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func BenchmarkLookupAll1(b *testing.B) {
	data, err := ioutil.ReadFile("testdata/05_maximum_01.in.data")
	if err != nil {
		b.Fatalf("read test data: %v", err)
	}
	target, err := ioutil.ReadFile("testdata/05_maximum_01.in.target")
	if err != nil {
		b.Fatalf("read test data: %v", err)
	}
	b.ResetTimer()

	sa := New(string(data))
	sa.Build()
	sa.LookupAll(string(target))
}

func BenchmarkLookupAll2(b *testing.B) {
	data, err := ioutil.ReadFile("testdata/05_maximum_01.in.data")
	if err != nil {
		b.Fatalf("read test data: %v", err)
	}
	target, err := ioutil.ReadFile("testdata/05_maximum_01.in.target")
	if err != nil {
		b.Fatalf("read test data: %v", err)
	}
	b.ResetTimer()

	index := suffixarray.New(data)
	index.Lookup(target, -1)
}

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
