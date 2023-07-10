package slicesx

func Filter[S ~[]T, T any](s S, f func(T) bool) (r S) {
	for i := range s {
		if f(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}

func Map[S ~[]T, M []N, T, N any](s S, f func(T) N) (m M) {
	m = make(M, len(s))
	for i := range s {
		m[i] = f(s[i])
	}
	return m
}

func Delete[S ~[]T, T any](s S, i int) (r S) {
	if i < 0 {
		return deleteByIndex(s, len(s)+i)
	}
	return deleteByIndex(s, i)
}

func deleteByIndex[S ~[]T, T any](s S, i int) (r S) {
	switch i {
	case 0:
		r = s[1:]
	case len(s) - 1:
		r = s[:i]
	default:
		r = make(S, len(s)-1)
		copy(r, s[:i])
		copy(r[i:], s[i+1:])
	}
	return
}
