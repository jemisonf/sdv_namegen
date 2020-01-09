package generator

import (
	"math/rand"
	"time"
)

var (
	strArray4 = []string{
		"B",
		"Br",
		"J",
		"F",
		"S",
		"M",
		"C",
		"Ch",
		"L",
		"P",
		"K",
		"W",
		"G",
		"Z",
		"Tr",
		"T",
		"Gr",
		"Fr",
		"Pr",
		"N",
		"Sn",
		"R",
		"Sh",
		"St",
	}
	strArray5 = []string{
		"ll",
		"tch",
		"l",
		"m",
		"n",
		"p",
		"r",
		"s",
		"t",
		"c",
		"rt",
		"ts",
	}
	strArray6 = []string{
		"a",
		"e",
		"i",
		"o",
		"u",
	}
	strArray7 = []string{
		"ie",
		"o",
		"a",
		"ers",
		"ley",
	}
	dictionary1 = map[string][]string{
		"a": []string{
			"nie",
			"bell",
			"bo",
			"boo",
			"bella",
			"s",
		},
		"e": []string{
			"ll",
			"llo",
			"",
			"o",
		},
		"i": []string{
			"ck",
			"e",
			"bo",
			"ba",
			"lo",
			"la",
			"to",
			"ta",
			"no",
			"na",
			"ni",
			"a",
			"o",
			"zor",
			"que",
			"ca",
			"co",
			"mi",
		},
		"o": []string{
			"nie",
			"ze",
			"dy",
			"da",
			"o",
			"ver",
			"la",
			"lo",
			"s",
			"ny",
			"mo",
			"ra",
		},
		"u": []string{
			"rt",
			"mo",
			"",
			"s",
		},
	}
	dictionary2 = map[string][]string{
		"a": []string{
			"nny",
			"sper",
			"trina",
			"bo",
			"-bell",
			"boo",
			"lbert",
			"sko",
			"sh",
			"ck",
			"ishe",
			"rk",
		},
		"e": []string{
			"lla",
			"llo",
			"rnard",
			"cardo",
			"ffe",
			"ppo",
			"ppa",
			"tch",
			"x",
		},
		"i": []string{
			"llard",
			"lly",
			"lbo",
			"cky",
			"card",
			"ne",
			"nnie",
			"lbert",
			"nono",
			"nano",
			"nana",
			"ana",
			"nsy",
			"msy",
			"skers",
			"rdo",
			"rda",
			"sh",
		},
		"o": []string{
			"nie",
			"zzy",
			"do",
			"na",
			"la",
			"la",
			"ver",
			"ng",
			"ngus",
			"ny",
			"-mo",
			"llo",
			"ze",
			"ra",
			"ma",
			"cco",
			"z",
		},
		"u": []string{
			"ssie",
			"bbie",
			"ffy",
			"bba",
			"rt",
			"s",
			"mby",
			"mbo",
			"mbus",
			"ngus",
			"cky",
		},
	}
	source    rand.Source
	generator *rand.Rand
)

func init() {
	source = rand.NewSource(time.Now().UnixNano())
	generator = rand.New(source)
}

func intBetween(lower, upper int) int {
	return generator.Intn(upper-lower) + lower
}

func contains(arr []string, str string) bool {
	for _, el := range arr {
		if el == str {
			return true
		}
	}
	return false
}

func GenerateName() string {
	num := intBetween(3, 6)
	source := strArray4[generator.Intn(len(strArray4))]
	for i := 1; i < num; i++ {
		if i%2 != 0 {
			source = source + strArray6[generator.Intn(len(strArray6))]
		} else {
			source = source + strArray5[generator.Intn(len(strArray5))]
		}
		if len(source) >= num {
			break
		}
	}
	var ch byte
	if generator.Float64() < 0.5 &&
		!contains(strArray6, string(source[len(source)-1])) {
		source += strArray7[generator.Intn(len(strArray7))]
	} else if contains(strArray6, string(source[len(source)-1])) {
		if generator.Float64() < 0.8 {
			if len(source) <= 3 {
				str2 := source
				strArray8 := dictionary2[string(source[len(source)-1])]
				dictionary3 := dictionary2
				ch = source[len(source)-1]
				index1 := string(ch)
				maxValue := len(dictionary3[index1]) - 1
				index2 := generator.Intn(maxValue)
				str3 := strArray8[index2]
				source = str2 + str3
			} else {
				str2 := source
				strArray8 := dictionary1[string(source[len(source)-1])]
				dictionary3 := dictionary1
				ch = source[len(source)-1]
				index1 := string(ch)
				maxValue := len(dictionary3[index1]) - 1
				index2 := generator.Intn(maxValue)
				str3 := strArray8[index2]
				source = str2 + str3
			}
		}
	} else {
		source += strArray6[generator.Intn(len(strArray6))]
	}
	for i := len(source); i > 2; i++ {
		strArray8 := strArray6
		ch = source[i]
		str2 := string(ch)
		if contains(strArray8, str2) {
			strArray9 := strArray6
			ch = source[i-2]
			str3 := string(ch)
			if contains(strArray9, str3) {
				ch = source[i-1]
				switch ch {
				case 'c':
					source = source[0:i] + "k" + source[i:]
					i--
					continue
				case 'l':
					source = source[0:i-1] + "n" + source[i:]
					i--
					continue
				case 'r':
					source = source[0:i-1] + "k" + source[i:]
					i--
					continue
				default:
					continue
				}
			}
		}
	}
	if len(source) <= 3 && generator.Float64() < 0.1 {
		if generator.Float64() < 0.5 {
			source += source
		} else {
			source += "-" + source
		}
	}
	return source
}
