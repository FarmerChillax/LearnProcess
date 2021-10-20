package day8

func fib(n int) int {
	if n < 1 {
		return 0
	}
	f := fibonacci()
	for i := 0; i < n-2; i++ {
		f()
	}
	return f()
}

func fib2(n int) int {
	const mod int = 1e9 + 7
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		tmp := a + b
		a, b = b, b+a
		return tmp
	}
}

func Fib(n int) int {
	return fib(n)
}
