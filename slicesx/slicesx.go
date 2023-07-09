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
	for i := range s {
		m = append(m, f(s[i]))
	}
	return m
}
