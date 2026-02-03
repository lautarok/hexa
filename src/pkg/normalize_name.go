package pkg

import "strings"

func NormalizeName(input string) string {

	input = strings.TrimSpace(input)
	input = strings.Join(strings.Fields(input), " ")

	words := strings.Split(input, " ")
	lowerParticles := map[string]bool{
		"de":  true,
		"da":  true,
		"do":  true,
		"van": true,
		"von": true,
		"der": true,
		"la":  true,
		"le":  true,
		"del": true,
	}

	for i, word := range words {
		lowerWord := strings.ToLower(word)

		if i > 0 && lowerParticles[lowerWord] {
			words[i] = lowerWord
			continue
		}

		words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	}

	return strings.Join(words, " ")
}
