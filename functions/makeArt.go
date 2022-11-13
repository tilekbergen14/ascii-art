package functions

func MakeArt(value string, data []byte) [][]string {
	symbolStartingPoint := DefineAsciiSymbol()
	asciiWordArr := [][]string{}
	for i := 0; i < len(value); i++ {
		subArr := []string{}
		count := 1
		word := ""
		if string(value[i]) == "\\" && string(value[i+1]) == "n" {
			subArr = append(subArr, "\n")
			i = i + 1
		} else if string(value[i]) == "\n" {
			subArr = append(subArr, "\n")
		} else {
			for _, a_char := range data {
				if a_char == '\n' {
					count++
				}
				if count >= symbolStartingPoint[int(value[i])] && count < symbolStartingPoint[int(value[i])]+9 {
					if a_char == '\n' {
						subArr = append(subArr, word)
						word = ""
					} else {
						word += string(a_char)
					}
				}
			}
		}
		asciiWordArr = append(asciiWordArr, subArr)
	}
	return asciiWordArr
}
