# Suffix Array

[![CircleCI](https://circleci.com/gh/d-tsuji/suffixarray.svg?style=svg)](https://app.circleci.com/pipelines/github/d-tsuji/suffixarray) [![GoDoc](https://godoc.org/github.com/d-tsuji/suffixarray?status.svg)](https://godoc.org/github.com/d-tsuji/suffixarray)

Manber&Myers's Suffix Array implemented in Go. I'm referring to [Manber.java](https://algs4.cs.princeton.edu/63suffix/Manber.java.html).

-  [U.Manber, G.Myers: "Suffix arrays: a new method for on-line string
    searches", SIAM Journal of Computing](https://karczmarczuk.users.greyc.fr/TEACH/TAL/Doc/BK_search/suffixAr.pdf)

## Verified

This library has been verified with the following problem.

- http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_14_D

## Benchmark

Here are the benchmark results.

BenchmarkLookupAll1 is my implementation of Manber&Myers' Algorithmic SuffixArray search. BenchmarkLookupAll2 is an official index/suffixarray search It is.

The implementation of index/suffixarray is faster on the order of about 10x.

```
> go test -bench .
goos: windows
goarch: amd64
pkg: github.com/d-tsuji/suffixarray
BenchmarkLookupAll1-8           1000000000               0.437 ns/op
BenchmarkLookupAll2-8           1000000000               0.0480 ns/op
PASS
ok      github.com/d-tsuji/suffixarray  10.743s
```

Because the SuffixArray construction of index/suffixarray uses the SA-IS algorithm, where N is the length of the string to be searched. It is O(N). On the other hand, Manber&Myers algorithm is O(N log(N)^2), where N is the length of the string to be searched for in large the effect of log(N)^2 increases as.

## LICENSE

These codes are licensed under CC0.

[![CC0](http://i.creativecommons.org/p/zero/1.0/88x31.png "CC0")](https://creativecommons.org/publicdomain/zero/1.0/)
