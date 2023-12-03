package main

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func rotate(s []int, diff int) []int {
	for i := 0; i < gcd(len(s), diff); i++ {
		for prev, next := i, i+diff; ; prev, next = next, prev+diff {
			if next >= len(s) {
				next %= len(s)
			}
			if next == i {
				break
			}
			s[next], s[prev] = s[prev], s[next]
		}
	}
	return s
}
