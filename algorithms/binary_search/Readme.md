### Binary Search

Finding first index for which predicate is true:


```go
// FirstTrue returns first index such that predicate is true
// if no such element NoMatch error is returned
func FirstTrue(x []int, p Pred) (int, error) {
	l := 0
	r := len(x) - 1

	for l < r {
		m := l + (r-l)/2
		if p(x[m]) {
			r = m
		} else {
			l = m + 1
		}
	}
	if p(x[l]) {
		return l, nil
	}
	return 0, NoMatch
}
```


Finding last index for which predicate is false: 

```go
// LastFalse returns last index such that predicate is true
// if no such element NoMatch error is returned
func LastFalse(x []int, p Pred) (int, error) {
	l := 0
	r := len(x) - 1
	for l < r {
		m := l + (r-l+1)/2
		if p(x[m]) {
			r = m - 1
		} else {
			l = m
		}
	}
	if !p(x[l]) {
		return l, nil
	}
	return 0, NoMatch
}


```

Link: https://www.topcoder.com/community/data-science/data-science-tutorials/binary-search/
