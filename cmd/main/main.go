package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/d-tsuji/suffixarray"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var s string
	var q int

	fmt.Fscan(r, &s)
	fmt.Fscan(r, &q)

	index := suffixarray.New(s)

	for i := 0; i < q; i++ {
		var p string
		fmt.Fscan(r, &p)

		res := index.LookupAll(p)
		if len(res) > 0 {
			fmt.Fprintln(w, 1)
		} else {
			fmt.Fprintln(w, 0)
		}
	}
}
