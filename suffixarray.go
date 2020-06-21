package suffixarray

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
