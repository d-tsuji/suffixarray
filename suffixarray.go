package suffixarray

import (
	"strings"
)

type Manber struct {
	// length of input string
	N int

	// input text (ASCII only)
	Text string

	// offset of ith string in order
	Index []int

	// Rank of ith string
	Rank []int

	// Rank of ith string (temporary)
	newrank []int

	offset int
}

// New creates a new Manber.
func New(s string) *Manber {
	n := len(s)
	m := &Manber{
		N:       n,
		Text:    s,
		Index:   make([]int, n+1),
		Rank:    make([]int, n+1),
		newrank: make([]int, n+1),
	}
	// sentinels
	m.Index[n] = n
	m.Rank[n] = -1

	return m
}

// Build builds a SuffixArray.
// Building time is O(N (logN)^2) where N is the
// size of the input string data.
func (m *Manber) Build() {
	m.msd()
	m.doit()
}

func (m *Manber) LookupAll(p string) []int {
	var left, right int

	// Find the maximum index where the result of strings.Compare is -1.
	l := 0
	r := m.N
	for r-l > 1 {
		mid := (l + r) >> 1
		cmp := strings.Compare(m.Text[m.Index[mid]:min(m.Index[mid]+len(p), m.N)], p)
		if cmp < 0 {
			l = mid
		} else {
			r = mid
		}
	}
	left = l

	// Find the maximum index where the result of strings.Compare is 0.
	l = 0
	r = m.N
	for r-l > 1 {
		mid := (l + r) >> 1
		cmp := strings.Compare(m.Text[m.Index[mid]:min(m.Index[mid]+len(p), m.N)], p)
		if cmp <= 0 {
			l = mid
		} else {
			r = mid
		}
	}
	right = l

	result := make([]int, 0, right-left)
	for i := left + 1; i <= right; i++ {
		result = append(result, m.Index[i])
	}
	return result
}

func (m *Manber) msd() {
	const R int = 256

	// calculate frequencies
	freq := make([]int, R)
	for i := 0; i < m.N; i++ {
		freq[m.Text[i]]++
	}

	// calculate cumulative frequencies
	cumm := make([]int, R)
	for i := 1; i < R; i++ {
		cumm[i] = cumm[i-1] + freq[i-1]
	}

	// compute ranks
	for i := 0; i < m.N; i++ {
		m.Rank[i] = cumm[m.Text[i]]
	}

	// sort by first char
	for i := 0; i < m.N; i++ {
		m.Index[cumm[m.Text[i]]] = i
		cumm[m.Text[i]]++
	}
}

func (m *Manber) doit() {
	for m.offset = 1; m.offset < m.N; m.offset += m.offset {
		var count int
		for i := 1; i <= m.N; i++ {
			if m.Rank[m.Index[i]] == m.Rank[m.Index[i-1]] {
				count++
			} else if count > 0 {
				// sort
				left := i - 1 - count
				right := i - 1
				m.quicksort(left, right)

				// now fix up ranks
				r := m.Rank[m.Index[left]]
				for j := left + 1; j <= right; j++ {
					if m.less(m.Index[j-1], m.Index[j]) {
						r = m.Rank[m.Index[left]] + j - left
					}
					m.newrank[m.Index[j]] = r
				}

				// copy back - note can't update rank too eagerly
				for j := left + 1; j <= right; j++ {
					m.Rank[m.Index[j]] = m.newrank[m.Index[j]]
				}

				count = 0
			}
		}
	}
}

// -----------------------------------------
// Helper functions for comparing suffixes.
// -----------------------------------------

func (m *Manber) quicksort(lo, hi int) {
	if hi <= lo {
		return
	}
	i := m.partition(lo, hi)
	m.quicksort(lo, i-1)
	m.quicksort(i+1, hi)
}

func (m *Manber) partition(lo, hi int) int {
	i, j, v := lo-1, hi, m.Index[hi]
	for {
		// find item on left to swap
		i++
		for m.less(m.Index[i], v) {
			if i == hi {
				break
			}
			i++
		}

		// find item on right to swap
		j--
		for m.less(v, m.Index[j]) {
			if j == lo {
				break
			}
			j--
		}

		// check if pointers cross
		if i >= j {
			break
		}
		m.exch(i, j)
	}

	// swap with partition element
	m.exch(i, hi)

	return i
}

func (m *Manber) exch(i, j int) {
	m.Index[i], m.Index[j] = m.Index[j], m.Index[i]
}

func (m *Manber) less(v, w int) bool {
	return m.Rank[v+m.offset] < m.Rank[w+m.offset]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
