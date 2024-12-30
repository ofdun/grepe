package io

var colors = map[string]string{
	"black":              "\033[30m",
	"red":                "\033[31m",
	"green":              "\033[32m",
	"yellow":             "\033[33m",
	"blue":               "\033[34m",
	"magenta":            "\033[35m",
	"cyan":               "\033[36m",
	"white":              "\033[37m",
	"black-background":   "\033[40m",
	"red-background":     "\033[41m",
	"green-background":   "\033[42m",
	"yellow-background":  "\033[43m",
	"blue-background":    "\033[44m",
	"magenta-background": "\033[45m",
	"cyan-background":    "\033[46m",
	"white-background":   "\033[47m",
}

func ColorExists(color string) bool {
	_, ok := colors[color]
	return ok
}

func getColor(color string) string {
	return colors[color]
}

func PrintPatternMatchesColorful(text *[]string, matches *[][][]int, color string, onlyMatching bool) {
	color = getColor(color)
	reset := "\033[0m"

	for i, str := range *text {
		curMatches := (*matches)[i]
		prev := 0
		if !onlyMatching {
			for _, pair := range curMatches {
				print(str[prev:pair[0]], color, str[pair[0]:pair[1]], reset)
				prev = pair[1]
			}
			println(str[prev:])
		} else {
			for _, pair := range curMatches {
				println(color + str[pair[0]:pair[1]] + reset)
			}
		}
	}
}
