package io

func PrintPatternMatchesColorful(text *[]string, matches *[][][]int) {
	green := "\033[32m"
	reset := "\033[0m"

	for i, str := range *text {
		curMatches := (*matches)[i]
		prev := 0
		for _, pair := range curMatches {
			print(str[prev:pair[0]], green, str[pair[0]:pair[1]], reset)
			prev = pair[1]
		}
		println(str[prev:])
	}
}
