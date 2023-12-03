package main

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func totalPopCount(bytes []byte) int {
	total := 0
	for _, b := range bytes {
		total += int(pc[b])
	}
	return total
}
