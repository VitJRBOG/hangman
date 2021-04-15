package opponent

import (
	"strings"

	"bitbucket.org/VitJRBOG/hangman/tools"
)

// GetWord "загадывает" слово
func GetWord() string {
	path := tools.GetPath("words.txt")
	tools.CheckAndCreateFile(path)

	wordsLine := tools.ReadTextFile(path)
	if len(wordsLine) > 0 {
		words := convertLineToList(wordsLine)
		wordNumber := randomizeNumber()
		word := selectWord(wordNumber, words)
		return word
	}

	return ""
}

// convertLineToList преобразует строку со словами в слайс слов
func convertLineToList(wordsLine string) []string {
	return strings.Split(wordsLine, "\n")
}

// randomizeNumber генерирует случайное число
func randomizeNumber() int {
	// TODO: описать функцию
	return 0
}

// selectWord выбирает слово из списка по индексу
func selectWord(wordNumber int, words []string) string {
	return words[wordNumber]
}
