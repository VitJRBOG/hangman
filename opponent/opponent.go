package opponent

import (
	"math/rand"
	"strings"
	"time"

	"github.com/VitJRBOG/hangman/tools"
)

// GetWord "загадывает" слово
func GetWord() string {
	path := tools.GetPath("words.txt")
	tools.CheckAndCreateFile(path)
	wordsLine := tools.ReadTextFile(path)
	if len(wordsLine) > 0 {
		word := selectWord(wordsLine)
		word = strings.ToLower(word)
		return word
	}
	return ""
}

// selectWord выбирает слово из списка
func selectWord(wordsLine string) string {
	words := convertLineToList(wordsLine)
	wordNumber := randomizeNumber(len(words))
	return words[wordNumber]
}

// convertLineToList преобразует строку со словами в слайс слов
func convertLineToList(wordsLine string) []string {
	return strings.Split(wordsLine, "\n")
}

// randomizeNumber возвращает псевдослучайное число от 0 до n
func randomizeNumber(n int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(n)
	return number
}
