package judgeman

import "strings"

// CheckLetter проверят наличие буквы в слове
func CheckLetter(letter, word string) []int {
	letters := strings.Split(word, "")
	var indexes []int
	for i := 0; i < len(letters); i++ {
		if letters[i] == letter {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// CheckWord проверяет законченность слова
func CheckWord(wordField string) bool {
	if x := strings.Index(wordField, "_"); x == -1 {
		return true
	}
	return false
}

// CheckHangmanImage проверяет законченность рисунка висельника
func CheckHangmanImage(hangmanImage string) bool {
	fullHangmanImage := "\n" +
		"  ------------\n" +
		" ||          |\n" +
		" ||         { }\n" +
		" ||        / | \\\n" +
		" ||       /  |  \\\n" +
		" ||         / \\\n" +
		" ||        /   \\\n" +
		" /\\       /     \\\n" +
		"/  \\\n"

	return hangmanImage == fullHangmanImage
}
