package gopurs_strings_extra

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func isWordChar(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsSymbol(r) || unicode.Is(unicode.Sc, r) || unicode.Is(unicode.So, r) || unicode.Is(unicode.Lm, r)
}

func Words(str string) []string {
	var words []string
	var current []rune

	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if !isWordChar(r) {
			if len(current) > 0 {
				words = append(words, string(current))
				current = nil
			}
			continue
		}

		if i > 0 && len(current) > 0 {
			prev := current[len(current)-1]
			// lowercase to uppercase transition
			if unicode.IsLower(prev) && unicode.IsUpper(r) {
				words = append(words, string(current))
				current = nil
			}
		}
		current = append(current, r)
	}
	if len(current) > 0 {
		words = append(words, string(current))
	}
	if words == nil {
		words = []string{} // Return empty array, not nil for purescript
	}
	return words
}

func CamelCase(str string) string {
	w := Words(str)
	if len(w) == 0 {
		return ""
	}
	res := strings.ToLower(w[0])
	for i := 1; i < len(w); i++ {
		runes := []rune(strings.ToLower(w[i]))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		res += string(runes)
	}
	return res
}

func PascalCase(str string) string {
	w := Words(str)
	if len(w) == 0 {
		return ""
	}
	var res string
	for i := 0; i < len(w); i++ {
		runes := []rune(strings.ToLower(w[i]))
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		res += string(runes)
	}
	return res
}

func KebabCase(str string) string {
	w := Words(str)
	var lowerWords []string
	for _, word := range w {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}
	return strings.Join(lowerWords, "-")
}

func SnakeCase(str string) string {
	w := Words(str)
	var lowerWords []string
	for _, word := range w {
		lowerWords = append(lowerWords, strings.ToLower(word))
	}
	return strings.Join(lowerWords, "_")
}

func UpperCaseFirst(str string) string {
	if len(str) == 0 {
		return str
	}
	r, size := utf8.DecodeRuneInString(str)
	return strings.ToUpper(string(r)) + strings.ToLower(str[size:])
}

func Levenshtein(s1 string, s2 string) int {
	r1 := []rune(s1)
	r2 := []rune(s2)

	len1 := len(r1)
	len2 := len(r2)

	if len1 == 0 {
		return len2
	}
	if len2 == 0 {
		return len1
	}

	d := make([][]int, len1+1)
	for i := range d {
		d[i] = make([]int, len2+1)
		d[i][0] = i
	}
	for j := range d[0] {
		d[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			cost := 1
			if r1[i-1] == r2[j-1] {
				cost = 0
			}

			min := d[i-1][j] + 1
			if d[i][j-1]+1 < min {
				min = d[i][j-1] + 1
			}
			if d[i-1][j-1]+cost < min {
				min = d[i-1][j-1] + cost
			}
			d[i][j] = min
		}
	}
	return d[len1][len2]
}

func SorensenDiceCoefficient(s1 string, s2 string) float64 {
	r1 := []rune(s1)
	r2 := []rune(s2)

	if len(r1) < 2 || len(r2) < 2 {
		return 0
	}

	bigrams := make(map[string]int)
	for i := 0; i < len(r1)-1; i++ {
		bg := string(r1[i : i+2])
		bigrams[bg]++
	}

	intersection := 0
	for i := 0; i < len(r2)-1; i++ {
		bg := string(r2[i : i+2])
		if count, ok := bigrams[bg]; ok && count > 0 {
			bigrams[bg]--
			intersection++
		}
	}

	return (2.0 * float64(intersection)) / float64(len(r1)+len(r2)-2)
}
