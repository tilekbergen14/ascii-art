package functions

func DefineAsciiSymbol() map[int]int {
	symbols := make(map[int]int)
	count := 2
	for i := 32; i < 127; i++ {
		symbols[i] = count
		count += 9
	}
	return symbols
}
