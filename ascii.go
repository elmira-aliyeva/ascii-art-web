package ascii

import (
	"io/ioutil"
	"os"
	"strings"
)

const height = 8

// Asciify is
func Asciify(args ...string) (string, error) {
	fonts := []string{"standard", "thinkertoy", "shadow"}
	var words, filename, saveTo string
	var ascii, output, fs = false, false, false
	for i, arg := range args {
		for index, font := range fonts {
			if (arg == font && i == len(args)-2 && args[len(args)-1][:9] == "--output=") || (arg == font && i == len(args)-1) {
				filename = arg + ".txt"
				if i == len(args)-2 {
					saveTo = args[len(args)-1][9:]
					output = true
				} else {
					fs = true
				}
			} else if i == len(args)-1 && index == len(fonts)-1 && !fs && !output {
				words += arg
				filename = "standard.txt"
				ascii = true
			}
		}
		if !ascii && !output && !fs {
			words += arg
			if i != len(args)-1 {
				words += " "
			}
		}
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	arrData := strings.Split(string(data), "\n")
	var m = make(map[rune][]string)
	var runeIt = ' '
	var first = 1
	for index, line := range arrData {
		if index >= first && index <= first+height {
			m[runeIt] = append(m[runeIt], line)
			if index == first+height {
				runeIt++
				first += height + 1
			}
		}
	}
	var file *os.File
	var err1 error
	var res string
	if output {
		file, err1 = os.Create(saveTo)
		if err1 != nil {
			return "", err1
		}
	}
	for _, set := range strings.Split(words, "\\n") {
		for line := 0; line < height; line++ {
			for _, runa := range set {
				if ascii || fs {
					res += printLine(m[runa][line], line)
				} else if output {
					saveArt(m[runa][line], line, file)
				}
			}
			if ascii || fs {
				// fmt.Print("\n")
				res += "\n"
			} else if output {
				file.WriteString("\n")
			}
		}
	}
	return res, err
}

func printLine(mapStr string, line int) string {
	var str string
	for _, symbol := range mapStr {
		str += string(symbol)
	}
	return str
}

func saveArt(mapStr string, line int, file *os.File) {
	for _, symbol := range mapStr {
		file.WriteString(string(symbol))
	}
}
