package main

import (
	"fmt"

	"bitbucket.org/VitJRBOG/hangman/opponent"
	"bitbucket.org/VitJRBOG/hangman/playground"
)

func main() {
	mistakesNumber := 0
	word := opponent.GetWord()
	wordField, hangmanImage := playground.PreparingPlayground(word)
	for {
		letter := playground.GetLetterFromUser()
		indexes := playground.LookingLetterInWord(letter, word)
		mistakesNumber, wordField, hangmanImage = playground.Painting(mistakesNumber, indexes, letter, wordField, hangmanImage)
		fmt.Printf("%s\n%s\n\n", hangmanImage, wordField)
		if needContinue := playground.Summarizing(hangmanImage, wordField); !needContinue {
			break
		}
	}
}
