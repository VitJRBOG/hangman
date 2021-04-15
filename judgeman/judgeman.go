package judgeman

import "strings"

// CheckLetter проверят наличие буквы в слове
func CheckLetter(letter, word string) []int {
	var indexes []int
	for i := 0; i < len(word); i++ {
		if x := strings.Index(word, letter); x != -1 {
			indexes = append(indexes, x)
		}
	}
	return indexes
}

// CheckWord проверяет законченность слова
func CheckWord(wordField string) bool {
	if x := strings.Index(wordField, "_"); x == -1 {
		return false
	}
	return true
}

// CheckHangmanImage проверяет законченность рисунка висельника
func CheckHangmanImage(hangmanImage string) bool {
	fullHangmanImage := `------------
                        ||          |
                        ||         { }
                        ||        / | \
                        ||       /  |  \
                        ||         / \
                        ||        /   \
                        /\       /     \
                       /  \`

	return hangmanImage == fullHangmanImage
}
