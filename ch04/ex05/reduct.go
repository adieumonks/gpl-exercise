package main

func reduct(s []string) []string {
	if len(s) == 0 {
		return s
	}
	j := 0
	for i := 0; i < len(s)-1; i++ {
		if s[j] != s[i+1] {
			s[j+1] = s[i+1]
			j++
		}
	}
	return s[:j+1]
}
